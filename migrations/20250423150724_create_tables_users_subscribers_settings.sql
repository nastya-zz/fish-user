-- +goose Up
-- +goose StatementBegin

-- Пользователи с RBAC
CREATE TABLE users (
    id UUID PRIMARY KEY UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    avatar_path VARCHAR(255),
    bio TEXT,
    experience_level INTEGER CHECK (experience_level BETWEEN 1 AND 10) DEFAULT 1,
    is_verified BOOLEAN DEFAULT FALSE,
    is_blocked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

-- 2. Таблица подписок (follows)

CREATE TABLE follows (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     follower_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
     following_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
     status VARCHAR(20) DEFAULT 'pending', -- 'pending', 'accepted', 'rejected' (для закрытых аккаунтов)
     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
     updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    -- Один пользователь может подписаться на другого только один раз
     CONSTRAINT unique_follow UNIQUE (follower_id, following_id),

    -- Проверка, что пользователь не подписывается сам на себя
     CONSTRAINT no_self_follow CHECK (follower_id != following_id)
);

-- Индексы для быстрого поиска
CREATE INDEX idx_follows_follower ON follows(follower_id);
CREATE INDEX idx_follows_following ON follows(following_id);


-- 3. Таблица запросов на подписку (для закрытых аккаунтов)
CREATE TABLE follow_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    requester_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    target_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'pending', -- 'pending', 'accepted', 'rejected'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    CONSTRAINT unique_follow_request UNIQUE (requester_id, target_id),
    CONSTRAINT no_self_request CHECK (requester_id != target_id)
    );

-- Индексы
CREATE INDEX idx_follow_requests_requester ON follow_requests(requester_id);
CREATE INDEX idx_follow_requests_target ON follow_requests(target_id);


-- 4. Таблица блокировок пользователей
CREATE TABLE user_blocks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    blocker_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    blocked_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    CONSTRAINT unique_block UNIQUE (blocker_id, blocked_id),
    CONSTRAINT no_self_block CHECK (blocker_id != blocked_id)
);

-- Индексы
CREATE INDEX idx_user_blocks_blocker ON user_blocks(blocker_id);
CREATE INDEX idx_user_blocks_blocked ON user_blocks(blocked_id);

-- Настройки
CREATE TYPE Language AS ENUM ('RU', 'ENG');
CREATE TYPE Availability AS ENUM ('PUBLIC', 'PRIVATE');
CREATE TABLE settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    language  Language NOT NULL DEFAULT 'RU',
    availability Availability NOT NULL DEFAULT 'PUBLIC'
);

CREATE TABLE user_settings (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    settings_id UUID REFERENCES settings(id),
    PRIMARY KEY (user_id, settings_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

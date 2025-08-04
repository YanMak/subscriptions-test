-- Индекс для фильтрации по user_id
CREATE INDEX idx_user_id ON subscriptions(user_id);

-- Композитный индекс для ускорения обновлений по id + user_id
CREATE INDEX idx_id_user_id ON subscriptions(id, user_id);
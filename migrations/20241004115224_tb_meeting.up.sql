CREATE TABLE IF NOT EXISTS tb_meeting (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Идентификатор встречи
    name VARCHAR(100) NOT NULL, -- Название
    place VARCHAR(100), -- Место
    comment VARCHAR(1000), -- Комментарий
    applicant_id VARCHAR(128) NOT NULL, -- Соискатель
    start_date DATE NOT NULL, -- Дата начала
    end_date DATE NOT NULL, -- Дата окончания
    is_full_day BOOLEAN DEFAULT false, -- Весь день
    is_online BOOLEAN DEFAULT false, -- Онлайн
    is_outlook_event BOOLEAN DEFAULT false, -- Создать событие в Outlook
    email_author VARCHAR(255) NOT NULL, -- Email организатора
    is_view_meeting BOOLEAN DEFAULT false, -- Признак просмотра уведомления за 15 минут
    is_start_meeting BOOLEAN DEFAULT false -- Признак просмотра уведомления после начала    
);
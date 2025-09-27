-- Habilita a verificação de chaves estrangeiras para garantir a integridade dos dados.
PRAGMA foreign_keys = ON;

-- =================================================================================
-- ESTRUTURA DAS TABELAS (FASE 1)
-- =================================================================================

-- Tabela 1
-- A tabela mais alta da nossa aplicação. É através dela que chegamos em qualquer dado.
CREATE TABLE simulation_runs (
    id INTEGER NOT NULL PRIMARY KEY,
    run_name TEXT,
    number_of_houses INTEGER NOT NULL,
    number_of_days INTEGER NOT NULL,
    number_of_residents INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    -- No futuro, você pode adicionar uma coluna aqui:
    -- simulation_profile_id INTEGER
);

-- NOVA TABELA: Armazena os participantes estáticos de uma simulação
CREATE TABLE run_participants (
    id INTEGER NOT NULL PRIMARY KEY,                -- O RG único do registro no DB
    run_id INTEGER NOT NULL,                        -- A qual simulação pertence
    resident_id_in_run INTEGER NOT NULL,         -- O ID GLOBAL do participante (1 até N total)
    house_id INTEGER NOT NULL,                      -- A qual casa o participante pertence
    resident_id_in_house INTEGER NOT NULL,          -- O ID LOCAL do participante DENTRO da casa (1 até M)
    house_profile_id INTEGER,
    resident_profile_id INTEGER,
    age INTEGER,
    FOREIGN KEY (run_id) REFERENCES simulation_runs(id) ON DELETE CASCADE
);



-- Tabela de logs diários agora é muito mais simples
CREATE TABLE daily_logs (
    id INTEGER NOT NULL PRIMARY KEY,
    participant_id INTEGER NOT NULL, -- Liga ao participante na tabela acima
    day INTEGER NOT NULL,
    FOREIGN KEY (participant_id) REFERENCES run_participants(id) ON DELETE CASCADE
);

-- Tabela 2b: Eventos de Uso Sanitário
-- Esta é a tabela mais detalhada. Cada linha é um único evento de uso.
-- Ela armazena as informações dos seus structs `Sanitary` e `Usage`.
CREATE TABLE usage_events (
    id INTEGER NOT NULL PRIMARY KEY,
    daily_log_id INTEGER NOT NULL,                -- Liga este evento ao residente/dia específico na tabela acima
    sanitary_type TEXT NOT NULL,                  -- "toilet", "shower", "washBassin", etc.
    sanitary_device_id INTEGER,               -- Corresponde a `sanitaryDeviceID`
    start_usage INTEGER NOT NULL,                 -- Corresponde a `startUsage` (segundo do dia)
    end_usage INTEGER NOT NULL,                   -- Corresponde a `endUsage` (segundo do dia)
    flow_rate REAL NOT NULL,                      -- Corresponde a `flowRate`

    FOREIGN KEY (daily_log_id) REFERENCES daily_logs(id) ON DELETE CASCADE
);


CREATE TABLE TiposResidencia (
    id INT UNSIGNED PRIMARY KEY,    -- Seu HouseID original. Não precisa de AUTO_INCREMENT se os IDs são pré-definidos.
    descricao VARCHAR(255) NOT NULL -- Ex: "Casa Popular", "Apartamento 2 quartos", "Kitnet"
);

CREATE TABLE Ocupacoes (
    id INT UNSIGNED PRIMARY KEY,   -- Seu ResidentOccupationID original.
    nome VARCHAR(100) NOT NULL UNIQUE -- Ex: "Estudante", "Trabalhador", "Aposentado"
);

CREATE TABLE PerfisDeMorador (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tipo_residencia_id INT UNSIGNED NOT NULL,
    ocupacao_id INT UNSIGNED NOT NULL,
    idade TINYINT UNSIGNED NOT NULL,

    -- Garante que não existam perfis duplicados. Essencial para a integridade!
    UNIQUE KEY uk_perfil (tipo_residencia_id, ocupacao_id, idade),

    FOREIGN KEY (tipo_residencia_id) REFERENCES TiposResidencia(id),
    FOREIGN KEY (ocupacao_id) REFERENCES Ocupacoes(id)
);

CREATE TABLE CategoriasDispositivos (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL UNIQUE -- "Chuveiro", "Vaso Sanitário", "Pia de Cozinha"
);

CREATE TABLE ModelosDispositivos (
    id INT UNSIGNED PRIMARY KEY,              -- Seu SanitaryDeviceID original.
    categoria_id INT UNSIGNED NOT NULL,       -- Chave estrangeira para a categoria geral.
    descricao VARCHAR(255),                   -- Ex: "Chuveiro Elétrico 5500W", "Vaso com Caixa Acoplada 6L"

    FOREIGN KEY (categoria_id) REFERENCES CategoriasDispositivos(id)
);

CREATE TABLE Simulacoes (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao TEXT,
    data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Você pode adicionar outros metadados aqui, como:
    -- parametros_json TEXT     -- Para guardar os parâmetros de entrada da simulação
    UNIQUE KEY uk_nome (nome) -- Garante que cada simulação tenha um nome único
);

CREATE TABLE LogsDeUso (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    simulacao_id INT UNSIGNED NOT NULL,            -- <-- A NOVA COLUNA!
    dia SMALLINT UNSIGNED NOT NULL,
    perfil_morador_id INT UNSIGNED NOT NULL,
    modelo_dispositivo_id INT UNSIGNED NOT NULL,
    inicio_uso_segundos INT NOT NULL,
    fim_uso_segundos INT NOT NULL,
    vazao_ls DOUBLE PRECISION NOT NULL,

    FOREIGN KEY (simulacao_id) REFERENCES Simulacoes(id) ON DELETE CASCADE,
    FOREIGN KEY (perfil_morador_id) REFERENCES PerfisDeMorador(id),
    FOREIGN KEY (modelo_dispositivo_id) REFERENCES ModelosDispositivos(id)
);
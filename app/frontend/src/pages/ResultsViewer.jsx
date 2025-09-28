import { useState } from 'react';
// Importe a função que acabamos de criar no backend Go
import { SelectFile } from '../../wailsjs/go/main/App';

function ResultsViewer() {
    // State para gerenciar o arquivo e o feedback da UI
    const [fileName, setFileName] = useState(null);
    const [error, setError] = useState(null);
    const [isDragOver, setIsDragOver] = useState(false);

    // Função que processa um arquivo, seja do clique ou do drop
    const processFile = (file) => {
        // Validação: verifica se o arquivo existe e se é um .csv
        if (!file || !file.name.toLowerCase().endsWith('.csv')) {
            setError('Erro: Por favor, selecione apenas arquivos .csv.');
            setFileName(null);
            return;
        }

        setError(null);
        setFileName(file.name);
        
        // --- PRÓXIMO PASSO: ENVIAR PARA O BACKEND PARA PARSE ---
        // A propriedade 'path' está disponível em arquivos via drag-and-drop no Wails
        const filePath = file.path; 
        console.log("Arquivo válido selecionado:", filePath);
        // Aqui você chamaria outra função do Go para ler e processar o CSV
        // Ex: ParseCSV(filePath).then(data => setChartData(data));
    };

    // Função chamada pelo botão "Selecionar Arquivo"
    const handleFileSelectClick = () => {
        setError(null);
        SelectFile()
            .then(filePath => {
                // Se o usuário selecionou um arquivo (não cancelou)
                if (filePath) {
                    // A função do Go já garante que é um .csv
                    const name = filePath.split(/[\\/]/).pop(); // Pega apenas o nome do arquivo do caminho completo
                    setFileName(name);
                    console.log("Arquivo válido selecionado via diálogo:", filePath);
                    // Aqui você chamaria a função de parse do backend com o filePath
                }
            })
            .catch(err => {
                setError(`Erro ao abrir diálogo: ${err}`);
            });
    };

    // --- FUNÇÕES DE DRAG AND DROP ---

    const handleDragOver = (e) => {
        e.preventDefault(); // Necessário para permitir o drop
        e.stopPropagation();
        setIsDragOver(true);
    };

    const handleDragLeave = (e) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragOver(false);
    };

    const handleDrop = (e) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragOver(false);

        // Pega o primeiro arquivo que foi solto
        const droppedFile = e.dataTransfer.files && e.dataTransfer.files[0];
        processFile(droppedFile);
    };

    return (
        <div className="page-container">
            <h3>Visualizador de Resultados</h3>
            <p>Selecione um arquivo de análise (.csv) ou arraste e solte na área abaixo para visualizar os gráficos de consumo.</p>
            
            <button className="btn" onClick={handleFileSelectClick}>
                Selecionar Arquivo de Análise
            </button>
            
            {/* Área de Drag and Drop */}
            <div 
                className={`drop-zone ${isDragOver ? 'drag-over' : ''}`}
                onDragOver={handleDragOver}
                onDragLeave={handleDragLeave}
                onDrop={handleDrop}
            >
                {error ? (
                    <p className="drop-zone-error">{error}</p>
                ) : fileName ? (
                    <p className="drop-zone-success">Arquivo selecionado: <strong>{fileName}</strong></p>
                ) : (
                    <p>Arraste e solte um arquivo .csv aqui</p>
                )}
            </div>
            
            <div className="chart-placeholder">
                <p>A área do gráfico aparecerá aqui.</p>
            </div>
        </div>
    );
}

export default ResultsViewer;
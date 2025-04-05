function dom = uso_lavatorio(dom,dias)

%% Usos do lavatório

%  Sorteia a frequencia e vazão de uso do lavatório por morador / dia
    N = dom.nmoradores;
    for i = 1: N
        for j = 1: dias
            % sorteia a frequencia de uso
            freq = ceil(random('Poisson',5.93));
            dom.morador(i).lavatorio(j).frequencia = freq;
            
            %  Sorteia a vazão    
            if freq == 0
                
                vazao   = 0;
                duracao = 0;
                
            else
        
               [duracao, vazao]=lavatorio_function(freq);
               

                duracoes = dom.lavatorio(i).duracao(j).dia;
                duracoes = incrementa(duracao,duracoes);
                dom.lavatorio(i).duracao(j).dia = duracoes;
                        
                vazoes   = dom.lavatorio(i).vazao(j).dia;            
                vazoes   = incrementa(vazao,vazoes);
                
                dom.lavatorio(i).vazao(j).dia   = vazoes;
                dom.lavatorio(i).consumo(j).dia   = sum(dom.lavatorio(i).vazao(j).dia.*dom.lavatorio(i).duracao(j).dia);
                dom = hor_lavatorio(dom,dias,i,j,duracao); 
 %{           
                dom.morador(i).lavatorio(j).duracao = duracoes;
                dom.morador(i).lavatorio(j).vazao   = vazoes;
                dom.morador(i).lavatorio(j).consumo = consumo;
 %}           
                   
                         
        end
        
        end
  
end 
end
function dom = hor_lavatorio(dom,dias,i,j,duracao)  
% j -  dias de análise
% i - pessoa da análise
 %% Definição dos horários de uso do lavatório     

            % Definição dos horários das atividades
            pessoa = dom.pessoas(i);
            [time]= horario_function(dias,pessoa,dom);
           
            
            % Definição dos horários de uso do lavatório              
            freq       = dom.morador(i).lavatorio(j).frequencia;            
            [hora_lav] = sorteio_hor_lavatorio(time,j,freq,duracao,dom);
            
            % Atualiza os horários de uso do lavatório
       %     dom.morador(i).lavatorio(j).horario=hora_lav;

            horas = dom.lavatorio(i).horario(j).dia;
            horas = incrementa(hora_lav,horas);          
            dom.lavatorio(i).horario(j).dia = duration(0,0,horas,'Format','dd:hh:mm:ss');  
            
end

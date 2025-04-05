function dom = uso_maquina(dom,dias)

%% Usos da m�quina

%  Sorteia a frequencia e vaz�o de uso da maquina por morador / dia
    N = dom.nmoradores;
    for j = 1: dias
       for i = 1: N
            % sorteia a frequencia de uso
            freq = ceil(random('Poisson',0.37));
            dom.maquina(j).frequencia = freq;
            
            %  Sorteia a vaz�o    
            if freq == 0
                
                vazao   = [];
                duracao = [];
                dom.maquina(j).horario=[];
            else
        
               [duracao, vazao]=maquina_function(freq);
               dom = hor_maquina(dom,dias,i,j,duracao);

            end   
                dom.maquina(j).duracao   = duracao;            
                dom.maquina(j).vazao   = vazao;
                dom.maquina(j).consumo = sum(dom.maquina(j).vazao.*dom.maquina(j).duracao);
                dom = hor_maquina(dom,dias,i,j,duracao); 
        
                                                            
               
              
        end
        
    end
end


function dom = hor_maquina(dom,dias,i,j,duracao)
% j -  dias de an�lise
% i - pessoa da an�lise
%% Defini��o dos hor�rios de uso da m�quina de lavar 
            
            % Defini��o dos hor�rios das atividades
            pessoa = dom.pessoas(i);
            [time]= horario_function(dias,pessoa);            
                      
            freq      = dom.maquina(j).frequencia;  
            if freq==0
                dom.maquina(j).horario=[];
            else
            [hora_maq]= sorteio_hor_maquina(time,pessoa,j,freq,duracao);
            
            % Atualiza os hor�rios de uso da maquina
            dom.maquina(j).horario = duration(0,0,hora_maq,'Format','dd:hh:mm:ss');
            end
          
end

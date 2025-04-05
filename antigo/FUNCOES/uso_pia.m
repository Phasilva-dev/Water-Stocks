function dom = uso_pia(dom,dias)


%% Usos do lavatório

%  Sorteia a frequencia e vazão de uso da pia por morador / dia
    N = dom.nmoradores;
    for j = 1: dias
       for i = 1: N
            % sorteia a frequencia de uso
            freq = ceil(random('Poisson',24.88));
            dom.pia_cozinha(j).frequencia = freq;
            
            %  Sorteia a vazão    
            if freq == 0
                
                vazao   = 0;
                duracao = 0;
                
            else
        
               [duracao, vazao]=pia_cozinha_function(freq);
               dom = hor_pia(dom,dias,i,j,duracao);

            end
                dom.pia_cozinha(j).duracao = duracao;             
            
                dom.pia_cozinha(j).vazao = vazao;
                
                dom.pia_cozinha(j).consumo = sum(dom.pia_cozinha(j).duracao.*dom.pia_cozinha(j).vazao);
               
            
             
            
               
                                          
              
        end
        
    end
end



function dom = hor_pia(dom,dias,i,j,duracao)
% j -  dias de análise
% i - pessoa da análise            
%% Definição dos horários de uso da pia sanitária 

            % Definição dos horários das atividades
            pessoa = dom.pessoas(i);
            [time]= horario_function(dias,pessoa);
            

            freq      = dom.pia_cozinha(j).frequencia;            
            [hora_pia]= sorteio_hor_pia(time,j,freq,duracao);
            
            % Atualiza os horários de uso do lavatório
            dom.pia_cozinha(j).horario = duration(0,0,hora_pia,'Format','dd:hh:mm:ss');
        
end
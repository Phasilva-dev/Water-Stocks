function dom = uso_tanque(dom,dias)


%% Usos do tanque

%  Sorteia a frequencia e vazão de uso do tanque por morador / dia
    N = dom.nmoradores;
    for j = 1: dias
       for i = 1: N
            % sorteia a frequencia de uso
            freq = ceil(random('Poisson',1.15));
            dom.tanque(j).frequencia = freq;
            
            %  Sorteia a vazão    
            if freq == 0
                
                vazao   = 0;
                duracao = 0;
                
            else
        
               [duracao, vazao]=tanque_function(freq);
               
            end                
 
            dom.tanque(j).duracao = duracao;                                  
            dom.tanque(j).vazao   = vazao;
            dom.tanque(j).consumo = sum(dom.tanque(j).vazao.* dom.tanque(j).duracao);
            dom = hor_tanque(dom,dias,i,j);                
 

        end
        
    end
end


function dom = hor_tanque(dom,dias,i,j)
% j -  dias de análise
% i - pessoa da análise
%% Definição dos horários de uso do tanque
            
            % Definição dos horários das atividades
            pessoa = dom.pessoas(i);
            [time]= horario_function(dias,pessoa);            
                      

            freq       = dom.tanque(j).frequencia;            
            [hora_tanq]= sorteio_hor_tanque(time,j,freq);
            
            % Atualiza os horários de uso do Tanque
            dom.tanque(j).horario = duration(0,0,hora_tanq,'Format','dd:hh:mm:ss');
            
end
 

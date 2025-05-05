// Package behavioral fornece tipos e funções para definir perfis
// relacionados ao comportamento base de um residente como frequências e rotinas diárias,
// para simulações ou geração de dados.
package behavioral

import (
	"simulation/internal/dists" // Assumindo pacote local/terceiro com a interface Distribution.
	"errors"
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral" // Assumindo pacote local/terceiro com Routine e NewRoutine.
)

// RoutineProfile define um perfil de rotina diária para um residente.
// Ele contém uma sequência de distribuições estatísticas (`events`) que
// representam os horários de eventos (onde cada par geralmente significa
// acordar/sair ou entrar/dormir) e um `shift` que define um intervalo
// mínimo de tempo que deve ocorrer entre certos eventos consecutivos
// (por exemplo, entre acordar e sair de casa).
type RoutineProfile struct {
	// events contém pares de distribuições. Cada par representa um
	// ciclo de atividade, como (acordar, sair) ou (entrar, dormir).
	events []dists.Distribution // Não exportado. Use Events() para acesso.
	// shift define o intervalo mínimo (em unidades de tempo consistentes
	// com as distribuições) entre o primeiro e o segundo evento de um par
	// (ex: tempo mínimo acordado antes de sair).
	shift int32 // Não exportado. Use Shift() para acesso.
}

// Events retorna o slice interno de distribuições (`dists.Distribution`) que definem os
// horários dos eventos da rotina.
//
// ATENÇÃO: Conforme mencionado no código original, este getter fere o encapsulamento
// retornando uma referência direta ao slice interno. Isso foi feito intencionalmente
// para evitar a sobrecarga de cópia em cenários de uso muito intensivo (milhares de
// execuções). Tenha CUIDADO, pois modificações no slice retornado por este método
// afetarão diretamente o estado interno do RoutineProfile.
func (p *RoutineProfile) Events() []dists.Distribution {
	return p.events
}

// Shift retorna o valor do intervalo mínimo (`shift`) configurado para o perfil
// de rotina. Este valor representa a duração mínima entre certos eventos
// consecutivos (ex: tempo acordado antes de sair).
func (p *RoutineProfile) Shift() int32 {
	return p.shift
}

// NewRoutineProfile cria e inicializa uma nova instância de RoutineProfile.
//
// Parâmetros:
//   - shift: O intervalo mínimo (int32, deve ser não negativo) a ser aplicado
//     entre certos pares de eventos.
//   - events: Um slice de `dists.Distribution`. Deve conter um número par e
//     não zero de elementos, e nenhuma distribuição pode ser `nil`. Cada par
//     representa um ciclo de evento (ex: acordar/sair, entrar/dormir).
//
// Retorna:
//   - Um ponteiro (*RoutineProfile) para a instância criada e um erro `nil`.
//   - `nil` e um erro se as validações dos parâmetros falharem (events vazio,
//     tamanho ímpar, elemento nil, ou shift negativo).
//
// Importante: A função cria uma cópia interna do slice `events` fornecido para
// garantir que modificações externas no slice original não afetem o perfil criado.
func NewRoutineProfile(shift int32, events []dists.Distribution) (*RoutineProfile, error) {
	if len(events) == 0 {
		return nil, errors.New("o slice de eventos não pode estar vazio")
	}
	if len(events)%2 != 0 {
		return nil, errors.New("o número de elementos em eventos deve ser par")
	}
	if shift < 0 {
		// Considerar se 0 é um valor válido ou se deve ser > 0.
		// O código original permite 0, então mantemos.
		return nil, errors.New("o valor de shift (intervalo mínimo) deve ser não negativo")
	}

	// Criar uma cópia para garantir imutabilidade interna.
	eventsCopy := make([]dists.Distribution, len(events))
	copy(eventsCopy, events)

	// Validar se nenhuma distribuição na cópia é nula.
	for _, dist := range eventsCopy {
		if dist == nil {
			return nil, errors.New("nenhuma distribuição no slice de eventos pode ser nula")
		}
	}

	return &RoutineProfile{
		events: eventsCopy,
		shift:  shift,
	}, nil
}

// generateTime é uma função auxiliar (não exportada) que amostra um valor de tempo
// (float64) a partir da distribuição `dist` fornecida, usando o gerador `rng`,
// e o trunca para um valor `int32`.
func generateTime(dist dists.Distribution, rng *rand.Rand) int32 {
	var time float64 = dist.Sample(rng)
	// Simplesmente trunca a parte decimal.
	truncatedTime := int32(time)
	return truncatedTime
}

// enforceMinimunGap é um método auxiliar (não exportado) que garante que o
// intervalo entre um tempo de "entrada" (`entryTime`, ex: acordar) e o tempo
// de "saída" subsequente (`exitTime`, ex: sair de casa) seja de pelo menos `p.shift`.
//
// Se `exitTime` for menor que `entryTime + p.shift`, ele é ajustado para ser
// exatamente `entryTime + p.shift`. Caso contrário, o `exitTime` original é retornado.
func (p *RoutineProfile) enforceMinimunGap(entryTime, exitTime int32) int32 {
	// Se a diferença entre a saída e a entrada + shift for menor que o próprio shift
	// (ou seja, se exitTime < entryTime + shift), ajusta a saída.
	if exitTime-(entryTime+p.shift) < 0 { // Simplificado de `exitTime - (entryTime + p.shift) < 0`
		return entryTime + p.shift
	}
	return exitTime
}

// GenerateData gera uma instância completa de dados de rotina (`*residentdata.Routine`)
// com base nas distribuições de eventos e no `shift` (intervalo mínimo) do perfil.
//
// Utiliza o gerador de números aleatórios `rng` fornecido.
//
// O processo envolve:
// 1. Gerar um tempo bruto para cada evento no slice `p.events` usando `generateTime`.
// 2. Iterar sobre os eventos de índice ímpar (representando saídas ou o início do sono).
// 3. Para cada evento de índice ímpar, aplicar a regra de intervalo mínimo (`enforceMinimunGap`)
//    usando o evento de índice par imediatamente anterior (entrada ou acordar) e o tempo
//    gerado para o evento ímpar. Isso garante que haja pelo menos `p.shift` unidades de tempo
//    entre acordar e sair, ou entre chegar em casa e sair novamente (se aplicável ao modelo).
// 4. Construir e retornar um `*residentdata.Routine` com a sequência final de tempos ajustados.
//
// A lógica pressupõe que os eventos vêm em pares: o índice par é uma "entrada"
// (acordar, chegar em casa) e o índice ímpar seguinte é uma "saída" (sair de casa, dormir).
// O `shift` garante um tempo mínimo de permanência ou preparação.
func (p *RoutineProfile) GenerateData(rng *rand.Rand) *behavioral.Routine {
	// Aloca diretamente o slice para os tempos com o tamanho necessário.
	times := make([]int32, len(p.events))

	// 1. Gera tempos brutos para todos os eventos.
	for i, dist := range p.events {
		times[i] = generateTime(dist, rng)
	}

	// 2. Aplica a regra do intervalo mínimo (`shift`) aos eventos de índice ímpar.
	//    Itera sobre os índices ímpares (1, 3, 5, ...).
	for i := 1; i < len(p.events); i += 2 {
		// `times[i-1]` é o tempo de "entrada" (par).
		// `times[i]` é o tempo de "saída" (ímpar) bruto.
		// Ajusta `times[i]` (saída) para garantir o gap mínimo após `times[i-1]` (entrada).
		times[i] = p.enforceMinimunGap(times[i-1], times[i])
	}

	// 3. Cria a estrutura Routine com os tempos finais ajustados.
	return behavioral.NewRoutine(times)
}
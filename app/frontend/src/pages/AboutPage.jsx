// frontend/src/components/AboutPage.jsx

function AboutPage() {
  return (
    <div className="page-container about-page">
      <h3>Sobre o Simulador de Demanda de Água</h3>
      <p>
        O <strong>Simulador de Demanda de Água</strong> é uma ferramenta de simulação estocástica
        voltada para a estimativa do consumo residencial em sistemas de abastecimento.
        Desenvolvido como uma aplicação desktop com <strong>Go</strong> e <strong>React</strong> via <strong>Wails</strong>,
        o projeto oferece a engenheiros sanitaristas e planejadores urbanos uma abordagem robusta
        e flexível para analisar e prever padrões de uso da água.
      </p>
      
      <h4>Tecnologias Utilizadas</h4>
      <ul>
        <li><strong>Backend:</strong> Go (Golang), responsável pelos cálculos de simulação de alta performance.</li>
        <li><strong>Frontend:</strong> React.js, proporcionando uma interface de usuário moderna, interativa e responsiva.</li>
        <li><strong>Framework:</strong> Wails v2, que integra backend em Go e frontend em React em uma aplicação desktop nativa.</li>
        <li><strong>Biblioteca Estatística:</strong> Gonum, utilizada para operações matemáticas e estatísticas avançadas.</li>
      </ul>

      <h4>Licenças</h4>
      <p>
        Este software utiliza bibliotecas de código aberto. Agradecemos aos autores e comunidades responsáveis:
      </p>
      <ul>
        <li>
          <strong>Wails</strong> — Licença{" "}
          <a href="https://opensource.org/licenses/MIT" target="_blank" rel="noopener noreferrer">MIT</a><br />
          Copyright (c) 2018–Presente, Lea Anthony
        </li>
        <li>
          <strong>Plotly.js</strong> — Licença{" "}
          <a href="https://opensource.org/licenses/MIT" target="_blank" rel="noopener noreferrer">MIT</a><br />
          Copyright (c) 2016–2024, Plotly Technologies Inc.
        </li>
        <li>
          <strong>Gonum</strong> — Licença{" "}
          <a href="https://opensource.org/licenses/BSD-3-Clause" target="_blank" rel="noopener noreferrer">BSD 3-Clause</a><br />
          Copyright ©2013, The Gonum Authors
        </li>
      </ul>
      <p>
        As licenças completas estão incluídas neste projeto.
      </p>

      <h4>Créditos</h4>
      <p>
        Produzido como projeto de Iniciação Científica pela Universidade Estadual de Feira de Santana.<br />
        Desenvolvido por <strong>Pedro Henrique de Araújo Silva</strong>.<br />
        Idealizado por{" "}
        <a href="http://lattes.cnpq.br/123456" target="_blank" rel="noopener noreferrer">
          Prof. Eduardo Henrique Borges Cohim Silva
        </a>{" "}
        e{" "}
        <a href="http://lattes.cnpq.br/654321" target="_blank" rel="noopener noreferrer">
          Prof. Anderson de Souza Matos Gadea
        </a>.
      </p>

      <h4>Contato</h4>
      <p>
        Entre em contato pelo e-mail:{" "}
        <a href="mailto:phasilva.academic@gmail.com?subject=Contato%20sobre%20o%20Simulador%20de%20Demanda%20de%20Água">
          phasilva.academic@gmail.com
        </a>
      </p>
      <p className="contact-note">
        <em>
          Esta é a versão 1.0 do software. Estou aberto a contato para sugestões de melhoria,
          propostas de novas funcionalidades, esclarecimento de dúvidas relacionadas ao projeto, bem como para
          colaborações acadêmicas ou iniciativas de pesquisa e desenvolvimento.
        </em>
      </p>
    </div>
  );
}

export default AboutPage;

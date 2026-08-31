const UNITS = [
  {
    id: 1,
    title: "O Campo da Psicologia Social",
    icon: "📚",
    color: "#f97316",
    description: "História, epistemologia e paradigmas da disciplina",
    lessons: [
      {
        title: "O que é Psicologia Social?",
        content: `<p>A <strong>Psicologia Social</strong> estuda como o contexto social influencia pensamentos, sentimentos e comportamentos de indivíduos, grupos e instituições.</p>
        <div class="highlight"><strong>Objetivo central:</strong> compreender a relação entre processos subjetivos e realidade social, articulando teoria, pesquisa e prática.</div>
        <p>Diferencia-se da sociologia pelo foco na dimensão psicológica da experiência social e da psicologia individual pelo enfoque nas relações sociais.</p>`
      },
      {
        title: "Raízes Históricas",
        content: `<p>A Psicologia Social nasce do diálogo entre <strong>filosofia</strong> (Hobbes, Rousseau, Comte) e as <strong>ciências naturais</strong> do século XIX.</p>
        <ul>
          <li><strong>1898:</strong> Norman Triplett — primeiro experimento (efeito do público no desempenho)</li>
          <li><strong>1908:</strong> McDougall e Ross publicam os primeiros livros-texto</li>
          <li><strong>1924:</strong> Allport define a área como estudo científico do comportamento social</li>
        </ul>
        <p>Farr (2012) mostra que a "psicologia social moderna" se constitui em tradições distintas nos EUA e na Europa.</p>`
      },
      {
        title: "Paradigmas em Psicologia Social",
        content: `<p><strong>Paradigma</strong> = conjunto de pressupostos sobre o objeto, método e finalidade da ciência.</p>
        <ul>
          <li><strong>Positivismo:</strong> neutralidade, mensuração, causalidade, generalização</li>
          <li><strong>Crítico/histórico:</strong> sociedade como totalidade, ideologia, compromisso</li>
          <li><strong>Construcionista:</strong> realidade social como construção compartilhada</li>
        </ul>
        <div class="highlight">Não existe um único paradigma: tradições americana e europeia divergem sobre o que é "científico" em Psicologia Social.</div>`
      },
      {
        title: "Epistemologia e Conhecimento",
        content: `<p>A disciplina exige compreender os <strong>pressupostos epistemológicos</strong> que orientam pesquisa e intervenção.</p>
        <p><strong>Objetivos do curso:</strong> analisar relações entre contexto social e subjetividade; interpretar relatórios científicos; localizar e avaliar fontes bibliográficas.</p>
        <p>Competências: identificar fenômenos psicossociais nas relações entre indivíduos, grupos e instituições.</p>`
      },
      {
        title: "Pesquisa e Prática Social",
        content: `<p>A Psicologia Social articula <strong>produção científica</strong> e <strong>prática profissional</strong> — clínica, comunitária, organizacional, políticas públicas.</p>
        <p>O psicólogo social atua compreendendo representações, normas, identidades e processos de exclusão/inclusão.</p>
        <div class="highlight">A qualidade das fontes (SciELO, PePSIC, bases UNIP) é critério essencial para o Trabalho de Levantamento Bibliográfico (TLB).</div>`
      }
    ],
    quiz: [
      { q: "A Psicologia Social estuda principalmente:", options: ["Apenas neurotransmissores cerebrais","Relações entre contexto social e subjetividade","Somente testes de QI individual","Exclusivamente terapia psicanalítica"], answer: 1, explain: "A Psicologia Social foca processos psicossociais nas relações entre indivíduos, grupos e instituições." },
      { q: "Norman Triplett (1898) é lembrado por:", options: ["Criar a teoria das representações sociais","Realizar o primeiro experimento de Psicologia Social","Fundar o construcionismo social europeu","Escrever A construção social da realidade"], answer: 1, explain: "Triplett investigou o efeito da presença de outros no desempenho — marco experimental da área." },
      { q: "O paradigma positivista enfatiza:", options: ["Compromisso político explícito do pesquisador","Neutralidade, mensuração e generalização","Realidade como pura invenção individual","Rejeição total de qualquer método"], answer: 1, explain: "O positivismo busca objetividade, mensuração e leis gerais do comportamento social." },
      { q: "Allport (1924) definiu Psicologia Social como:", options: ["Estudo filosófico da moral","Estudo científico do comportamento social","Análise exclusiva de grupos pequenos","Teoria das representações sociais"], answer: 1, explain: "Allport consolidou a definição experimentalista dominante nos EUA." },
      { q: "Epistemologia em Psicologia Social refere-se a:", options: ["Técnicas de cirurgia cerebral","Pressupostos sobre conhecimento e método","Apenas estatística descritiva básica","História da arte contemporânea"], answer: 1, explain: "Epistemologia trata dos fundamentos e critérios de validade do conhecimento científico." }
    ],
    matching: [
      { term: "Triplett", match: "Primeiro experimento social" },
      { term: "Positivismo", match: "Neutralidade e mensuração" },
      { term: "Allport", match: "Comportamento social científico" },
      { term: "Epistemologia", match: "Fundamentos do conhecimento" },
      { term: "Psicologia Social", match: "Contexto e subjetividade" }
    ],
    boss: {
      name: "Guardião do Campo",
      desc: "Prove que domina os fundamentos da Psicologia Social!",
      questions: [
        { q: "A Psicologia Social diferencia-se da sociologia por:", options: ["Ignorar completamente a cultura","Focar a dimensão psicológica da experiência social","Estudar apenas instituições formais","Não usar nenhum método empírico"], answer: 1, explain: "Enquanto a sociologia prioriza estruturas sociais, a Psicologia Social enfatiza processos subjetivos." },
        { q: "Farr (2012) analisa principalmente:", options: ["Raízes da psicologia social moderna","Neuroanatomia do córtex frontal","Técnicas de reabilitação motora","Farmacologia dos antipsicóticos"], answer: 0, explain: "Farr mapeia as origens e tradições da psicologia social moderna." },
        { q: "Competência central do curso inclui:", options: ["Prescrever medicamentos controlados","Analisar fenômenos psicossociais","Realizar cirurgias neurológicas","Ignorar contexto histórico-social"], answer: 1, explain: "O plano de ensino prioriza análise crítica de fenômenos psicossociais." },
        { q: "Paradigma construcionista entende a realidade social como:", options: ["Dado objetivo imutável e fixo","Construção compartilhada e historicamente situada","Puramente biológica e genética","Inacessível a qualquer investigação"], answer: 1, explain: "O construcionismo social enfatiza a produção social do conhecimento e da realidade." },
        { q: "Bases de dados como SciELO e PePSIC servem para:", options: ["Jogar videogames online","Localizar fontes científicas confiáveis","Comprar livros de ficção","Publicar redes sociais pessoais"], answer: 1, explain: "São repositórios essenciais para o levantamento bibliográfico acadêmico." }
      ]
    }
  },
  {
    id: 2,
    title: "Psicologia Social nos EUA",
    icon: "🇺🇸",
    color: "#3b82f6",
    description: "Positivismo, experimentação e Kurt Lewin",
    lessons: [
      {
        title: "Ciência e Filosofia nos EUA",
        content: `<p>A Psicologia Social americana consolidou-se com a <strong>separação entre ciência e filosofia</strong>, adotando o modelo das ciências naturais.</p>
        <p>Após a Segunda Guerra Mundial, a área expandiu-se com financiamento federal e foco em problemas sociais mensuráveis.</p>
        <div class="highlight">Predomínio do <strong>individualismo metodológico</strong>: explicar o social a partir do indivíduo e de variáveis psicológicas.</div>`
      },
      {
        title: "Modelos Teórico-Metodológicos",
        content: `<p>Abordagens dominantes nos EUA:</p>
        <ul>
          <li><strong>Behaviorismo social:</strong> aprendizagem, reforço, modelagem</li>
          <li><strong>Cognitivismo social:</strong> atribuições, esquemas, processamento de informação</li>
          <li><strong>Experimentação de laboratório:</strong> controle de variáveis, amostras, estatística</li>
        </ul>
        <p>Crítica europeia: o laboratório artificializa fenômenos complexos do cotidiano.</p>`
      },
      {
        title: "Indivíduo e Grupo",
        content: `<p>Tema central: como o <strong>grupo</strong> influencia atitudes, conformidade, tomada de decisão e identidade.</p>
        <ul>
          <li><strong>Conformidade (Asch):</strong> pressão do grupo sobre julgamentos perceptivos</li>
          <li><strong>Obediência (Milgram):</strong> autoridade e cumprimento de ordens</li>
          <li><strong>Efeito espectador (Darley & Latané):</strong> difusão de responsabilidade</li>
        </ul>
        <p>A relação indivíduo-coletividade é estudada como variáveis mensuráveis e generalizáveis.</p>`
      },
      {
        title: "Kurt Lewin e a Ciência Pragmática",
        content: `<p><strong>Kurt Lewin</strong> (1890-1947), exilado nazista nos EUA, é pai da Psicologia Social moderna americana.</p>
        <ul>
          <li><strong>Teoria do campo:</strong> comportamento = f(pessoa, ambiente)</li>
          <li><strong>Ação research:</strong> pesquisa aliada à transformação social</li>
          <li><strong>Dinâmica de grupos:</strong> liderança, clima, tomada de decisão</li>
        </ul>
        <div class="highlight">Lewin articulou rigor científico e relevância prática — influência duradoura em dinâmica de grupos (Mailhiot, 2013).</div>`
      },
      {
        title: "Legado Experimental Americano",
        content: `<p>A tradição americana produziu teorias de:</p>
        <ul>
          <li><strong>Dissonância cognitiva</strong> (Festinger)</li>
          <li><strong>Teoria da identidade social</strong> (Tajfel — europeu, mas muito difundido nos EUA)</li>
          <li><strong>Atribuição causal</strong> (Heider, Kelley)</li>
        </ul>
        <p>Limitações: etnocentrismo, amostras WEIRD (Western, Educated, Industrialized, Rich, Democratic), descontextualização.</p>`
      }
    ],
    quiz: [
      { q: "A Psicologia Social americana priorizou:", options: ["Método experimental e positivismo","Apenas análise filosófica continental","Estudo exclusivo de representações sociais","Rejeição total de estatística"], answer: 0, explain: "A tradição americana consolidou o paradigma experimental-positivista." },
      { q: "Kurt Lewin propôs que comportamento depende de:", options: ["Apenas traços de personalidade fixos","Pessoa e ambiente (teoria do campo)","Somente fatores genéticos herdados","Exclusivamente estruturas cerebrais"], answer: 1, explain: "B = f(P, A) — comportamento é função da pessoa no ambiente." },
      { q: "O experimento de Asch demonstrou:", options: ["Obediência à autoridade legítima","Conformidade com julgamento do grupo","Difusão de responsabilidade em multidões","Formação de representações sociais"], answer: 1, explain: "Asch mostrou que pessoas alteram respostas para se conformar ao grupo." },
      { q: "Action research (pesquisa-ação) foi proposta por:", options: ["Moscovici", "Kurt Lewin", "Berger e Luckmann", "Sílvia Lane"], answer: 1, explain: "Lewin defendia pesquisa que transforma a realidade estudada." },
      { q: "Milgram investigou principalmente:", options: ["Conformidade visual em grupo","Obediência a ordens de autoridade","Memória de curto prazo verbal","Representações sobre saúde mental"], answer: 1, explain: "Milgram estudou até onde pessoas obedeciam ordens de dar choques." }
    ],
    matching: [
      { term: "Kurt Lewin", match: "Teoria do campo" },
      { term: "Asch", match: "Conformidade de grupo" },
      { term: "Milgram", match: "Obediência à autoridade" },
      { term: "Action Research", match: "Pesquisa e transformação" },
      { term: "Positivismo", match: "Método experimental" }
    ],
    boss: {
      name: "Titã Experimental",
      desc: "Enfrente o chefe da tradição americana!",
      questions: [
        { q: "Individualismo metodológico significa:", options: ["Explicar o social pelo indivíduo","Negar existência de grupos","Estudar apenas sociedades inteiras","Rejeitar qualquer experimento"], answer: 0, explain: "A tradição americana frequentemente reduz fenômenos sociais a variáveis individuais." },
        { q: "Dinâmica de grupos, segundo Lewin, estuda:", options: ["Apenas neurônios motores","Interações, liderança e clima grupal","Exclusivamente memória episódica","Somente desenvolvimento infantil"], answer: 1, explain: "Lewin pioneirizou o estudo sistemático de processos grupais." },
        { q: "Festinger é associado à teoria da:", options: ["Dissonância cognitiva","Representação social","Ideologia marxista clássica","Neuroplasticidade cerebral"], answer: 0, explain: "Festinger propôs que inconsistências cognitivas geram desconforto e mudança." },
        { q: "Crítica europeia ao laboratório americano aponta:", options: ["Excesso de relevância social","Artificialização e descontextualização","Uso insuficiente de estatística","Falta total de experimentos"], answer: 1, explain: "Europeus criticam a desconexão entre experimentos e vida social real." },
        { q: "WEIRD refere-se a amostras:", options: ["Ocidentais, educadas e industrializadas","Exclusivamente de países africanos","Apenas de crianças em idade escolar","Somente de populações rurais isoladas"], answer: 0, explain: "Crítica contemporânea: generalizações a partir de amostras não representativas." }
      ]
    }
  },
  {
    id: 3,
    title: "Psicologia Social Europeia",
    icon: "🇪🇺",
    color: "#8b5cf6",
    description: "Crítica ao positivismo, construcionismo e representações sociais",
    lessons: [
      {
        title: "A Resposta Europeia",
        content: `<p>A Psicologia Social europeia surge como <strong>crítica ao paradigma positivista americano</strong>, recuperando dimensões históricas, culturais e ideológicas.</p>
        <p>Após 1968, intensificam-se debates sobre compromisso social, linguagem e cotidiano.</p>
        <div class="highlight">A Europa questiona a neutralidade valorativa e o reducionismo experimental da tradição americana.</div>`
      },
      {
        title: "Processos de Socialização",
        content: `<p><strong>Socialização</strong> = processo pelo qual indivíduos internalizam normas, valores e papéis sociais ao longo da vida.</p>
        <ul>
          <li><strong>Socialização primária:</strong> família, infância</li>
          <li><strong>Socialização secundária:</strong> escola, mídia, trabalho</li>
          <li><strong>Agência social:</strong> indivíduos também transformam a sociedade</li>
        </ul>
        <p>A perspectiva europeia enfatiza a historicidade e a mediação simbólica desses processos.</p>`
      },
      {
        title: "Construcionismo Social",
        content: `<p><strong>Berger e Luckmann</strong> (1966/2014): a realidade social é construída na interação cotidiana.</p>
        <ul>
          <li><strong>Externalização:</strong> produção de objetos sociais</li>
          <li><strong>Objetivação:</strong> instituições parecem naturais</li>
          <li><strong>Internalização:</strong> incorporação de papéis e normas</li>
        </ul>
        <div class="highlight">"A realidade é socialmente construída" — conhecimento de senso comum torna-se tradição e instituição.</div>`
      },
      {
        title: "Representações Sociais",
        content: `<p><strong>Serge Moscovici</strong> (1925-2014) criou a teoria das <strong>Representações Sociais</strong> na década de 1960.</p>
        <p>RS = formas de conhecimento populares que tornam o real familiar, permitindo comunicação e orientação prática.</p>
        <ul>
          <li><strong>Ancoragem:</strong> classificar o novo em categorias conhecidas</li>
          <li><strong>Objetivação:</strong> criar imagens concretas de abstrações</li>
        </ul>
        <p>Referências: Guareschi & Jovchelovitch (2013); Spink (1995).</p>`
      },
      {
        title: "Cotidiano, Linguagem e Cultura",
        content: `<p>A tradição europeia estuda o <strong>senso comum</strong>, a <strong>linguagem</strong> e os <strong>discursos</strong> como produção social da realidade.</p>
        <p>Metodologias qualitativas ganham destaque: entrevistas, análise de conteúdo, estudos de caso, etnografia.</p>
        <div class="highlight">Haguette (2013) e autores do Centro Edelstein enfatizam metodologias qualitativas na sociologia e psicologia social.</div>`
      }
    ],
    quiz: [
      { q: "A Psicologia Social europeia criticou principalmente:", options: ["Uso de qualquer linguagem escrita","Paradigma positivista americano","Estudo de grupos pequenos","Teoria do campo de Lewin"], answer: 1, explain: "A tradição europeia contesta o experimentalismo descontextualizado dos EUA." },
      { q: "Berger e Luckmann são autores de:", options: ["Representações Sociais","A construção social da realidade","Dinâmica e Gênese dos Grupos","O que é psicologia social"], answer: 1, explain: "Obra clássica do construcionismo social na sociologia do conhecimento." },
      { q: "Moscovici é o principal teórico das:", options: ["Representações Sociais","Funções executivas","Ondas cerebrais alfa","Técnicas de dessensibilização"], answer: 0, explain: "Moscovici fundou a teoria das representações sociais na Europa." },
      { q: "Ancoragem, em Representações Sociais, significa:", options: ["Rejeitar qualquer novidade","Classificar o novo em categorias conhecidas","Copiar exatamente a tradição americana","Eliminar toda subjetividade"], answer: 1, explain: "Ancoragem integra o desconhecido a esquemas familiares do senso comum." },
      { q: "Socialização secundária ocorre principalmente em:", options: ["Útero materno exclusivamente","Escola, mídia e trabalho","Primeiros meses de vida apenas","Processos puramente genéticos"], answer: 1, explain: "Após a socialização primária (família), instituições medeiam novos papéis." }
    ],
    matching: [
      { term: "Moscovici", match: "Representações Sociais" },
      { term: "Berger & Luckmann", match: "Construção da realidade" },
      { term: "Ancoragem", match: "Classificar o novo" },
      { term: "Socialização", match: "Internalização de normas" },
      { term: "Tradição europeia", match: "Crítica ao positivismo" }
    ],
    boss: {
      name: "Oráculo Europeu",
      desc: "Demonstre domínio da tradição europeia!",
      questions: [
        { q: "Objetivação, para Berger e Luckmann, é quando:", options: ["Instituições parecem naturais e dadas","Indivíduos rejeitam toda sociedade","Experimentos perdem validade estatística","Grupos deixam de existir"], answer: 0, explain: "Produtos humanos adquirem aparência de objetividade e coerção." },
        { q: "Representações Sociais tornam o real:", options: ["Completamente inacessível","Familiar e comunicável no cotidiano","Puramente biológico e instintivo","Exclusivamente matemático"], answer: 1, explain: "RS traduzem o desconhecido em formas compreensíveis para o grupo." },
        { q: "Metodologias qualitativas na Europa privilegiam:", options: ["Apenas experimentos de laboratório","Significados, discursos e contexto","Exclusivamente escalas de QI","Somente neuroimagem funcional"], answer: 1, explain: "A tradição europeia valoriza compreensão interpretativa do social." },
        { q: "Spink (1995) discute representações sociais na perspectiva do:", options: ["Conhecimento no cotidiano","Desenvolvimento motor infantil","Processamento visual occipital","Tratamento farmacológico"], answer: 0, explain: "Spink articula RS com a vida cotidiana e o senso comum." },
        { q: "Internalização, no construcionismo, significa:", options: ["Rejeitar toda norma social","Incorporar papéis e visões de mundo","Destruir instituições existentes","Isolar completamente o indivíduo"], answer: 1, explain: "O sujeito assimila a realidade social como parte de sua identidade." }
      ]
    }
  },
  {
    id: 4,
    title: "Psicologia Social no Brasil",
    icon: "🇧🇷",
    color: "#10b981",
    description: "Compromisso social, ideologia e prática do psicólogo",
    lessons: [
      {
        title: "História no Brasil",
        content: `<p>A Psicologia Social brasileira desenvolveu-se entre a influência <strong>norte-americana</strong> (pós-guerra) e a recepção crítica da <strong>tradição europeia</strong>.</p>
        <p>Referências: Rodrigues, Assmar & Jablonsky (2016); Torres & Neiva (2011); Streit & Jacques (2016).</p>
        <div class="highlight">O campo brasileiro articula pesquisa acadêmica, movimentos sociais e reforma psicológica.</div>`
      },
      {
        title: "Sílvia Lane e a Mudança de Paradigma",
        content: `<p><strong>Sílvia Lane</strong> (1934-2019) foi referência na transformação da Psicologia Social brasileira.</p>
        <p>Carone (2007) destaca seu papel na crítica ao modelo importado e na defesa de uma psicologia comprometida com a realidade nacional.</p>
        <ul>
          <li>Obra: <em>O que é psicologia social</em> (Lane, 2017)</li>
          <li>Ênfase em história, cultura e prática social</li>
        </ul>`
      },
      {
        title: "Compromisso Social da Psicologia",
        content: `<p>A Psicologia brasileira, especialmente após os anos 1980, debate o <strong>compromisso social</strong> da profissão.</p>
        <p>O psicólogo não é neutro: suas práticas respondem a demandas sociais, políticas e éticas.</p>
        <div class="highlight">O TLB do 2º bimestre exige pesquisa em perspectiva social <strong>não positivista</strong> sobre o trabalho do psicólogo.</div>`
      },
      {
        title: "Ideologia e Prática Profissional",
        content: `<p><strong>Ideologia</strong> = conjunto de representações que naturalizam relações sociais e orientam a ação.</p>
        <p>A Psicologia Social brasileira analisa como discursos psicológicos participam da manutenção ou transformação social.</p>
        <p>Autores: Jacó-Vilela & Sato (2012); Rivero (2008); Zanela et al. (2008) — Centro Edelstein.</p>`
      },
      {
        title: "Psicologia e Práticas Sociais",
        content: `<p>O psicólogo atua em contextos diversos: saúde, educação, comunidade, organizações, direitos humanos.</p>
        <ul>
          <li><strong>Clínica ampliada:</strong> além do consultório</li>
          <li><strong>Psicologia social comunitária</strong></li>
          <li><strong>Políticas públicas</strong> e participação social</li>
        </ul>
        <div class="highlight">Formação exige criticidade teórica, ética profissional e capacidade de pesquisa bibliográfica qualificada.</div>`
      }
    ],
    quiz: [
      { q: "Sílvia Lane é referência na Psicologia Social:", options: ["Exclusivamente experimental americana","Brasileira e crítica ao modelo importado","Puramente neurobiológica clínica","Somente organizacional empresarial"], answer: 1, explain: "Lane liderou a crítica ao paradigma importado e a defesa de uma psicologia social brasileira." },
      { q: "Compromisso social da Psicologia implica:", options: ["Neutralidade absoluta do profissional","Responsabilidade ética e política na prática","Rejeitar todo contato com comunidades","Apenas atendimento individual privado"], answer: 1, explain: "O psicólogo reconhece que sua prática tem implicações sociais e políticas." },
      { q: "O TLB do 2º bimestre deve adotar perspectiva:", options: ["Exclusivamente positivista","Social não positivista","Puramente biológica","Sem qualquer referencial teórico"], answer: 1, explain: "O plano de ensino exige levantamento em perspectiva social crítica." },
      { q: "Lane (2017) publicou a obra:", options: ["A construção social da realidade","O que é psicologia social","Representações Sociais","Dinâmica de Grupo e Sistemas"], answer: 1, explain: "Obra básica da bibliografia do curso, de Sílvia Lane." },
      { q: "Ideologia, no contexto da disciplina, refere-se a:", options: ["Apenas propaganda eleitoral","Representações que naturalizam relações sociais","Técnicas de relaxamento muscular","Testes projetivos gráficos"], answer: 1, explain: "Ideologia orienta a compreensão e a ação no mundo social." }
    ],
    matching: [
      { term: "Sílvia Lane", match: "Mudança de paradigma no Brasil" },
      { term: "Compromisso social", match: "Responsabilidade ético-política" },
      { term: "TLB", match: "Levantamento bibliográfico" },
      { term: "Ideologia", match: "Naturalização social" },
      { term: "Centro Edelstein", match: "Pesquisa em psicologia social" }
    ],
    boss: {
      name: "Guardião Brasileiro",
      desc: "O chefe final! Prove que domina todo o semestre!",
      questions: [
        { q: "Carone (2007) analisa o papel de Sílvia Lane na:", options: ["Neuropsicologia clínica","Mudança da Psicologia Social do Brasil","Psicofarmacologia hospitalar","Avaliação neuropsicológica infantil"], answer: 1, explain: "Artigo na Psicologia & Sociedade sobre a contribuição de Lane." },
        { q: "Streit & Jacques (2016) organizam obra sobre:", options: ["Psicologia social contemporânea","Neuroanatomia comparada","Estatística matemática avançada","Psicopatologia forense exclusiva"], answer: 0, explain: "Obra complementar da bibliografia: Psicologia social contemporânea." },
        { q: "A Psicologia Social brasileira articula:", options: ["Apenas importação de teorias","Pesquisa, cultura e práticas sociais","Somente experimentos de laboratório","Exclusivamente terapia cognitiva"], answer: 1, explain: "O campo nacional integra teoria importada com realidade e demandas locais." },
        { q: "Bases como PePSIC e SciELO são essenciais para:", options: ["Avaliar qualidade de fontes científicas","Jogar online com colegas","Substituir toda orientação docente","Evitar leitura de artigos"], answer: 0, explain: "Competência do curso: discriminar seriedade científica das fontes." },
        { q: "Prática social do psicólogo inclui:", options: ["Apenas consultório individual","Comunidade, políticas e direitos","Somente administração de testes","Exclusivamente prescrição médica"], answer: 1, explain: "A perspectiva social amplia os cenários de atuação profissional." }
      ]
    }
  }
];

const ACHIEVEMENTS = [
  { id: "first_lesson", icon: "📖", title: "Primeira Lição", desc: "Complete sua primeira lição", xp: 0 },
  { id: "first_quiz", icon: "✅", title: "Quizzer", desc: "Complete seu primeiro quiz", xp: 0 },
  { id: "first_boss", icon: "⚔️", title: "Caçador de Chefes", desc: "Derrote seu primeiro chefe", xp: 0 },
  { id: "unit1", icon: "📚", title: "Explorador do Campo", desc: "Complete a Unidade 1", xp: 0 },
  { id: "unit2", icon: "🇺🇸", title: "Cientista Social", desc: "Complete a Unidade 2", xp: 0 },
  { id: "unit4", icon: "🏆", title: "Mestre Social", desc: "Complete todas as 4 unidades", xp: 0 },
  { id: "streak5", icon: "🔥", title: "Em Chamas", desc: "Acerte 5 questões seguidas", xp: 0 },
  { id: "streak10", icon: "💎", title: "Imparável", desc: "Acerte 10 questões seguidas", xp: 0 },
  { id: "no_life_lost", icon: "🛡️", title: "Invencível", desc: "Complete um quiz sem perder vidas", xp: 0 },
  { id: "level5", icon: "⭐", title: "Veterano", desc: "Alcance o nível 5", xp: 0 },
  { id: "level10", icon: "👑", title: "Lenda Social", desc: "Alcance o nível 10", xp: 0 },
  { id: "match_master", icon: "🔗", title: "Conector", desc: "Complete um desafio de associação", xp: 0 }
];

function stripLessonContent(html) {
  return html.replace(/<[^>]+>/g, ' ').replace(/\s+/g, ' ').trim();
}

const FLASHCARDS = UNITS.flatMap(u =>
  u.lessons.map(l => ({
    front: l.title,
    back: stripLessonContent(l.content),
    unit: u.title
  }))
);

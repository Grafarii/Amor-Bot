const UNITS = [
  {
    id: 1,
    title: "Escolas de Pensamento",
    icon: "🏛️",
    color: "#14b8a6",
    description: "Diversidade teórica, ontologia e epistemologia em Psicologia",
    lessons: [
      {
        title: "Psicologia ou Psicologias?",
        content: `<p>A Psicologia é marcada pela <strong>diversidade teórica</strong> — diferentes escolas partem de pressupostos ontológicos e epistemológicos distintos.</p>
        <div class="highlight">Ontologia = o que se considera real. Epistemologia = como se produz conhecimento válido.</div>
        <p>Figueiredo (2012) e Kahhale (2011) mostram que não há uma única "Psicologia", mas matrizes de pensamento historicamente situadas.</p>`
      },
      {
        title: "Contexto Histórico-Social",
        content: `<p>Teorias psicológicas emergem de <strong>cenários históricos, culturais e filosóficos</strong> específicos.</p>
        <ul>
          <li>Século XIX: laboratório, positivismo, ciência experimental</li>
          <li>Início do XX: crise de paradigmas, duas guerras mundiais</li>
          <li>Pós-guerra: expansão clínica, cognitivismo, crítica social</li>
        </ul>
        <p>Jacó-Vilela et al. (2013) enfatizam rumos e percursos da história da psicologia.</p>`
      },
      {
        title: "Organização em Escolas",
        content: `<p>As ideias psicológicas organizam-se em <strong>escolas de pensamento</strong> com autores, métodos e objetos privilegiados.</p>
        <p>Competência do curso: caracterizar abordagens e comparar convergências/divergências entre teorias.</p>
        <div class="highlight">Não se trata de "escolher a teoria certa", mas compreender pressupostos e limites de cada matriz.</div>`
      },
      {
        title: "Bock: Psicologias Plural",
        content: `<p><strong>Bock</strong> apresenta a Psicologia como campo plural de saberes e práticas.</p>
        <p>Cada abordagem define diferentemente: objeto de estudo, método, relação teoria-prática e papel do psicólogo.</p>
        <p>Schultz & Schultz (2019) oferece panorama da psicologia moderna em perspectiva histórica.</p>`
      },
      {
        title: "Leitura e Pesquisa",
        content: `<p>A disciplina exige <strong>ler e interpretar textos</strong> clássicos e levantar bibliografia em bases como PePSIC, SciELO e CAPES.</p>
        <p>No 2º bimestre: trabalho em grupo (até 5 alunos) sobre uma abordagem teórica, em ABNT.</p>
        <div class="highlight">Objetivo: relacionar formulações teóricas aos contextos que as tornaram possíveis.</div>`
      }
    ],
    quiz: [
      { q: "A diversidade teórica em Psicologia decorre principalmente de:", options: ["Falta de universidades no mundo","Diferentes pressupostos ontológicos e epistemológicos","Ausência total de método científico","Unicidade de objeto de estudo"], answer: 1, explain: "Cada escola parte de concepções distintas de realidade e de como conhecer." },
      { q: "Ontologia, em filosofia da ciência, trata de:", options: ["Técnicas estatísticas de pesquisa","O que se considera real ou existente","Normas de formatação ABNT","Procedimentos de laboratório"], answer: 1, explain: "Ontologia questiona a natureza do ser e do real." },
      { q: "Figueiredo (2012) aborda principalmente:", options: ["Neuroanatomia do tálamo","Matrizes do pensamento psicológico","Técnicas de dessensibilização sistemática","Farmacologia dos antidepressivos"], answer: 1, explain: "Obra clássica da bibliografia sobre matrizes teóricas." },
      { q: "Competência central da disciplina inclui:", options: ["Prescrever psicofármacos","Caracterizar principais abordagens teóricas","Realizar cirurgias neurológicas","Ignorar contexto histórico"], answer: 1, explain: "O curso forma leitura crítica das matrizes do pensamento psicológico." },
      { q: "Teorias psicológicas são historicamente:", options: ["Situadas e culturalmente produzidas","Atemporais e universais sem contexto","Puramente biológicas e genéticas","Identicas em todas as culturas"], answer: 0, explain: "Contexto histórico-social condiciona surgimento e sentido das teorias." }
    ],
    matching: [
      { term: "Ontologia", match: "Natureza do real" },
      { term: "Epistemologia", match: "Como se conhece" },
      { term: "Figueiredo", match: "Matrizes do pensamento" },
      { term: "Escola de pensamento", match: "Conjunto teórico organizado" },
      { term: "Bock", match: "Pluralidade de psicologias" }
    ],
    boss: {
      name: "Guardião das Matrizes",
      desc: "Prove que domina os fundamentos das escolas de pensamento!",
      questions: [
        { q: "Kahhale (2011) discute principalmente:", options: ["Diversidade teórica da Psicologia","Neuroplasticidade cerebral","Estatística inferencial avançada","Psicopatologia forense"], answer: 0, explain: "A Diversidade da Psicologia: Uma Construção Teórica." },
        { q: "Comparar teorias implica identificar:", options: ["Apenas semelhanças superficiais","Convergências e divergências fundamentais","Somente datas de nascimento dos autores","Exclusivamente métodos estatísticos"], answer: 1, explain: "Análise comparativa examina pressupostos e diferenças conceituais." },
        { q: "Jacó-Vilela et al. (2013) abordam:", options: ["História da Psicologia","Anatomia do sistema límbico","Técnicas de relaxamento muscular","Administração hospitalar"], answer: 0, explain: "História da Psicologia: Rumos e Percursos." },
        { q: "Psicologia como saber científico é:", options: ["Referido a condições históricas e sociais","Isento de qualquer valor ou ideologia","Puramente experimental desde a Antiguidade","Identico à medicina neurológica"], answer: 0, explain: "Competência do curso: reconhecer historicidade da produção científica." },
        { q: "O trabalho do 2º bimestre exige pesquisa em:", options: ["Uma abordagem teórica em ABNT","Exclusivamente testes de QI","Somente artigos de neuroimagem","Apenas manuais de farmácia"], answer: 0, explain: "TLB sobre abordagem estudada no semestre, em grupo." }
      ]
    }
  },
  {
    id: 2,
    title: "Os Behaviorismos",
    icon: "🐀",
    color: "#3b82f6",
    description: "Watson, Skinner e o behaviorismo radical",
    lessons: [
      {
        title: "Behaviorismo de Watson",
        content: `<p><strong>John B. Watson</strong> (1878-1958) lançou o manifesto behaviorista em 1913, rompendo com introspecção e psicologia da consciência.</p>
        <div class="highlight">Tese central: Psicologia deve estudar <strong>comportamento observável</strong>, não estados mentais privados.</div>
        <p>Condição reflexa: estímulo → resposta. Aprendizagem por associação e condicionamento clássico (Pavlov).</p>`
      },
      {
        title: "Behaviorismo Radical",
        content: `<p><strong>B. F. Skinner</strong> desenvolve o behaviorismo radical: comportamento é função de histórias de reforço.</p>
        <ul>
          <li><strong>Condicionamento operante:</strong> consequências moldam comportamento</li>
          <li><strong>Reforço positivo/negativo:</strong> aumenta probabilidade de resposta</li>
          <li><strong>Punição:</strong> diminui comportamento (uso controverso)</li>
        </ul>
        <p>Skinner (2011): Sobre o Behaviorismo — obra básica da bibliografia.</p>`
      },
      {
        title: "Conceitos-Chave",
        content: `<p><strong>Estímulo discriminativo:</strong> sinal que indica probabilidade de reforço.</p>
        <p><strong>Extinção:</strong> cessar reforço reduz comportamento.</p>
        <p><strong>Generalização e discriminação:</strong> respostas se estendem ou restringem a estímulos similares.</p>
        <p>Carrara & Zilio (2016) organizam reflexões históricas e conceituais sobre behaviorismos.</p>`
      },
      {
        title: "Críticas ao Behaviorismo",
        content: `<p>Críticas: reducionismo, negligência de cognição e significado, uso em controle social (Walden Two).</p>
        <p>Contribuições: método experimental rigoroso, base da modificação comportamental e análise aplicada.</p>
        <div class="highlight">O behaviorismo influenciou TCC, terapia comportamental e pesquisa em aprendizagem.</div>`
      },
      {
        title: "Contexto Histórico",
        content: `<p>Surgimento ligado ao <strong>positivismo</strong> e ao ideal de ciência mensurável no início do século XX.</p>
        <p>Watson prometeu predizer e controlar comportamento — promessa alinhada ao espírito científico da época.</p>
        <p>Behaviorismo dominou psicologia americana entre décadas de 1920-1950.</p>`
      }
    ],
    quiz: [
      { q: "Watson defendeu que a Psicologia deve estudar:", options: ["Apenas inconsciente freudiano","Comportamento observável e mensurável","Exclusivamente fenômenos transcendentais","Somente estruturas cerebrais"], answer: 1, explain: "Watson rejeitou introspecção e defendeu comportamento público." },
      { q: "Skinner é associado ao:", options: ["Behaviorismo radical e operante","Psicanálise lacaniana","Gestalt de Wertheimer","Fenomenologia husserliana"], answer: 0, explain: "Skinner focou reforço e condicionamento operante." },
      { q: "Condicionamento operante relaciona comportamento a:", options: ["Consequências e reforços","Arquétipos junguianos","Insight gestáltico repentino","Análise existencial do ser"], answer: 0, explain: "Consequências ambientais moldam probabilidade de respostas." },
      { q: "O manifesto behaviorista de Watson é de:", options: ["1913","1898","1968","1859"], answer: 0, explain: "1913 marca a ruptura com psicologia da consciência." },
      { q: "Carrara & Zilio (2016) abordam:", options: ["Reflexões sobre behaviorismos","Teoria das representações sociais","Neuroanatomia comparada","Existencialismo sartreano"], answer: 0, explain: "Behaviorismos: reflexões históricas e conceituais." }
    ],
    matching: [
      { term: "Watson", match: "Comportamento observável" },
      { term: "Skinner", match: "Reforço operante" },
      { term: "Condicionamento clássico", match: "Estímulo-resposta associado" },
      { term: "Extinção", match: "Cessa reforço" },
      { term: "Behaviorismo radical", match: "Função de reforço" }
    ],
    boss: {
      name: "Titã Comportamental",
      desc: "Enfrente o chefe do behaviorismo!",
      questions: [
        { q: "Pavlov está ligado ao:", options: ["Condicionamento clássico","Análise existencial","Teoria do amadurecimento","Materialismo histórico"], answer: 0, explain: "Condicionamento clássico: associação estímulo neutro e incondicionado." },
        { q: "Skinner utilizou principalmente:", options: ["Caixa operante e reforço","Divã e associação livre","Experiência eidética","Análise de conteúdo social"], answer: 0, explain: "A caixa operante permitiu controlar consequências e medir respostas." },
        { q: "Crítica comum ao behaviorismo watsoniano:", options: ["Ignora cognição e subjetividade","Excesso de foco no inconsciente","Dependência de fenomenologia","Rejeição total de experimentos"], answer: 0, explain: "Behaviorismo clássico excluiu processos mentais mediadores." },
        { q: "Reforço negativo consiste em:", options: ["Remover estímulo aversivo","Adicionar punição severa","Ignorar todo comportamento","Rejeitar método experimental"], answer: 0, explain: "Remover algo desagradável aumenta comportamento — ainda é reforço." },
        { q: "Behaviorismo influenciou diretamente:", options: ["Terapias comportamentais e TCC","Psicanálise kleiniana","Gestalt de Berlin","Marxismo vygotskiano"], answer: 0, explain: "Base empírica para modificação comportamental e cognitivo-comportamental." }
      ]
    }
  },
  {
    id: 3,
    title: "Funcionalismo Europeu",
    icon: "🔷",
    color: "#8b5cf6",
    description: "Gestalt e Kurt Lewin",
    lessons: [
      {
        title: "Psicologia da Gestalt",
        content: `<p><strong>Gestalt</strong> (forma/configuração): "O todo é diferente da soma das partes".</p>
        <p>Autores: Wertheimer, Köhler, Koffka — surgimento na Alemanha, década de 1910.</p>
        <div class="highlight">Percepção organiza-se em <strong>campos totais</strong> com leis: proximidade, similaridade, continuidade, fechamento.</div>`
      },
      {
        title: "Insight e Aprendizagem",
        content: `<p>Köhler estudou chimpanzés e o fenômeno do <strong>insight</strong> — compreensão súbita da solução.</p>
        <p>Crítica ao trial-and-error behaviorista: aprendizagem pode ser qualitativa, não apenas associação.</p>
        <p>Gestalt expandiu-se para psicoterapia (Perls, Gestalt-terapia) e cognição.</p>`
      },
      {
        title: "Kurt Lewin e Teoria de Campo",
        content: `<p><strong>Lewin</strong>: B = f(P, E) — comportamento é função da pessoa no ambiente.</p>
        <ul>
          <li><strong>Campo psicológico:</strong> totalidade de fatores presentes no momento</li>
          <li><strong>Dinâmica de grupos</strong> e action research</li>
          <li>Ponte entre experimentalismo e relevância social</li>
        </ul>`
      },
      {
        title: "Funcionalismo vs Estruturalismo",
        content: `<p>O funcionalismo europeu perguntava <strong>para que serve</strong> a consciência e o comportamento, não apenas sua estrutura.</p>
        <p>William James (EUA) também é referência funcionalista, em diálogo com pragmatismo.</p>
        <p>Na Europa, Gestalt e Lewin respondem à mecanicismo associationista.</p>`
      },
      {
        title: "Legado",
        content: `<p>Gestalt influenciou psicologia da percepção, cognição e terapias humanistas.</p>
        <p>Lewin influenciou psicologia social, organizacional e estudos de grupo.</p>
        <div class="highlight">Ambos enfatizam <strong>totalidade</strong> e contexto, antecipando críticas ao reducionismo.</div>`
      }
    ],
    quiz: [
      { q: "O princípio gestáltico afirma que:", options: ["O todo difere da soma das partes","A mente é tabula rasa ao nascer","Inconsciente determina tudo","Comportamento é só reflexo"], answer: 0, explain: "Organização perceptiva cria qualidades emergentes do campo." },
      { q: "Köhler é conhecido pelo estudo de:", options: ["Insight em chimpanzés","Condicionamento operante","Complexo de Édipo","Representações sociais"], answer: 0, explain: "Demonstrações de solução súbita de problemas." },
      { q: "Lewin propôs que comportamento depende de:", options: ["Pessoa e ambiente (campo)","Somente genes hereditários","Exclusivamente pulsões sexuais","Apenas reforços externos"], answer: 0, explain: "B = f(P, E) — teoria de campo." },
      { q: "Leis gestálticas de organização perceptiva incluem:", options: ["Proximidade e similaridade","Reforço e punição","Projeção e identificação","Sublimação e recalque"], answer: 0, explain: "Leis descrevem como percebemos formas unificadas." },
      { q: "Gestalt surgiu principalmente na:", options: ["Alemanha, início do séc. XX","Inglaterra vitoriana","Rússia czarista","China imperial"], answer: 0, explain: "Berlim, década de 1910, com Wertheimer e Köhler." }
    ],
    matching: [
      { term: "Gestalt", match: "Todo e forma" },
      { term: "Köhler", match: "Insight" },
      { term: "Lewin", match: "Teoria de campo" },
      { term: "Wertheimer", match: "Percepção de movimento" },
      { term: "B = f(P,E)", match: "Pessoa e ambiente" }
    ],
    boss: {
      name: "Mestre da Forma",
      desc: "Demonstre domínio do funcionalismo europeu!",
      questions: [
        { q: "Gestalt-terapia associa-se principalmente a:", options: ["Perls e consciência presente","Skinner e reforço","Freud e recalque","Beck e distorções cognitivas"], answer: 0, explain: "Fritz Perls desenvolveu abordagem gestáltica em terapia." },
        { q: "Funcionalismo pergunta principalmente:", options: ["Para que serve o psíquico","Qual peso do cérebro","Quantos neurônios existem","Qual QI médio populacional"], answer: 0, explain: "Foco na função adaptativa dos processos mentais." },
        { q: "Lewin contribuiu para estudos de:", options: ["Dinâmica de grupos","Neurotransmissores","Testes projetivos gráficos","Eletroencefalografia clínica"], answer: 0, explain: "Pioneiro na pesquisa de campo e grupos." },
        { q: "Gestalt critica o behaviorismo por:", options: ["Reduzir experiência a partes isoladas","Excesso de foco no inconsciente","Negar qualquer observação","Rejeitar experimentos de laboratório"], answer: 0, explain: "Gestalt defende organização total do campo experiencial." },
        { q: "William James representa funcionalismo:", options: ["Norte-americano pragmatista","Exclusivamente alemão","Puramente psicanalítico","Somente existencial"], answer: 0, explain: "James ligou função mental ao pragmatismo americano." }
      ]
    }
  },
  {
    id: 4,
    title: "A Psicanálise",
    icon: "🛋️",
    color: "#6366f1",
    description: "Freud e desdobramentos: ego, Klein e Winnicott",
    lessons: [
      {
        title: "Freud e a Criação da Psicanálise",
        content: `<p><strong>Sigmund Freud</strong> (1856-1939) funda a psicanálise em Viena, no final do século XIX.</p>
        <ul>
          <li><strong>Inconsciente:</strong> determinante do comportamento</li>
          <li><strong>Recalque, pulsões</strong> (Eros e Thanatos)</li>
          <li><strong>Desenvolvimento psicossexual:</strong> fases oral, anal, fálica, latência, genital</li>
        </ul>
        <div class="highlight">Método: associação livre, interpretação de sonhos, transferência.</div>`
      },
      {
        title: "Estrutura Tópica e Dinâmica",
        content: `<p>Modelo tópico: <strong>Isso, Ego e Superego</strong>.</p>
        <p>Modelo dinâmico: conflito entre pulsão, defesa e exigências da realidade.</p>
        <p>Mecanismos de defesa: negação, projeção, racionalização, sublimação.</p>`
      },
      {
        title: "Anna Freud e a Psicologia do Ego",
        content: `<p><strong>Anna Freud</strong> sistematizou a <strong>psicologia do ego</strong> e os mecanismos de defesa.</p>
        <p>Ênfase nas funções adaptativas do ego, não apenas nas pulsões.</p>
        <p>Influência na psicologia do ego norte-americana (Hartmann, Kris).</p>`
      },
      {
        title: "Melanie Klein e Relações Objetais",
        content: `<p><strong>Melanie Klein</strong> desenvolveu a <strong>teoria das relações objetais</strong>.</p>
        <p>Foco nas relações internas com objetos primários (mãe), posições paranoide-esquizóide e depressiva.</p>
        <p>Ansiedade primitiva e fantasia inconsciente desde a infância.</p>`
      },
      {
        title: "D. W. Winnicott",
        content: `<p><strong>Winnicott</strong> propôs a <strong>teoria do amadurecimento emocional</strong>.</p>
        <ul>
          <li><strong>Ambiente facilitador</strong> e mãe "suficientemente boa"</li>
          <li><strong>Objeto transicional</strong> (mantinha, etc.)</li>
          <li>Verdadeiro self vs falso self</li>
        </ul>
        <div class="highlight">Ênfase na holding, brincar e desenvolvimento em relação.</div>`
      }
    ],
    quiz: [
      { q: "Freud considerava central o conceito de:", options: ["Inconsciente e conflito psíquico","Reforço operante","Campo perceptual total","Materialismo histórico"], answer: 0, explain: "Inconsciente e conflito estruturam a teoria freudiana." },
      { q: "Anna Freud é associada à:", options: ["Psicologia do ego","Behaviorismo radical","Gestalt de percepção","TCC de Beck"], answer: 0, explain: "Sistematizou defesas do ego e sua adaptação." },
      { q: "Melanie Klein desenvolveu:", options: ["Teoria das relações objetais","Condicionamento clássico","Analítica existencial","Terapia racional emotiva"], answer: 0, explain: "Foco em objetos internos e posições esquizóide-depressiva." },
      { q: "Winnicott cunhou o conceito de:", options: ["Objeto transicional","Complexo de Édipo","Insight gestáltico","Dissonância cognitiva"], answer: 0, explain: "Objeto transicional media entre self e mundo." },
      { q: "O Superego freudiano relaciona-se a:", options: ["Moral, culpa e proibição","Percepção imediata","Reforço positivo","Campo visual"], answer: 0, explain: "Superego internaliza normas e exigências morais." }
    ],
    matching: [
      { term: "Freud", match: "Inconsciente e pulsão" },
      { term: "Anna Freud", match: "Psicologia do ego" },
      { term: "Klein", match: "Relações objetais" },
      { term: "Winnicott", match: "Amadurecimento emocional" },
      { term: "Transferência", match: "Revivescência na análise" }
    ],
    boss: {
      name: "Oráculo Freudiano",
      desc: "Enfrente o chefe da psicanálise!",
      questions: [
        { q: "Associação livre é técnica:", options: ["Psicanalítica freudiana","Behaviorista skinneriana","Humanista rogeriana","Gestalt perlsiana"], answer: 0, explain: "Paciente verbaliza sem censura para acessar inconsciente." },
        { q: "Recalque consiste em:", options: ["Excluir conteúdos da consciência","Reforçar comportamento desejado","Organizar percepção em gestalt","Racionalizar ações socialmente"], answer: 0, explain: "Defesa que empurra conteúdos inaceitáveis ao inconsciente." },
        { q: "Posição depressiva (Klein) envolve:", options: ["Reparação e integração do objeto","Apenas reforço negativo","Insight perceptual súbito","Análise de contingências"], answer: 0, explain: "Criança integra amor e ódio dirigidos ao objeto." },
        { q: "Mãe 'suficientemente boa' é conceito de:", options: ["Winnicott","Watson","Skinner","Wertheimer"], answer: 0, explain: "Ambiente que adapta-se às necessidades emergentes do bebê." },
        { q: "Psicanálise é frequentemente chamada de:", options: ["Primeira grande teoria clínica do séc. XX","Terceira força em psicologia","Única abordagem científica válida","Teoria puramente experimental"], answer: 0, explain: "Influência massiva na clínica, cultura e humanidades." }
      ]
    }
  },
  {
    id: 5,
    title: "Cognitivo-Comportamental",
    icon: "💭",
    color: "#06b6d4",
    description: "TCC, Beck e bases epistemológicas",
    lessons: [
      {
        title: "Surgimento da TCC",
        content: `<p>A <strong>Terapia Cognitivo-Comportamental (TCC)</strong> emerge na década de 1960 como síntese de tradições.</p>
        <p><strong>Beck</strong> (depressão, distorções cognitivas) e <strong>Ellis</strong> (REBT) são pilares.</p>
        <div class="highlight">Contexto: crise do behaviorismo radical por ignorar cognição; eficácia clínica mensurável.</div>`
      },
      {
        title: "Bases Epistemológicas",
        content: `<p>TCC combina:</p>
        <ul>
          <li><strong>Behaviorismo:</strong> comportamento observável, tarefas, experimentos comportamentais</li>
          <li><strong>Cognitivismo:</strong> crenças, esquemas, processamento de informação</li>
          <li><strong>Pragmatismo:</strong> foco em mudança e validação empírica</li>
        </ul>`
      },
      {
        title: "Modelo Teórico de Beck",
        content: `<p>Triade cognitiva na depressão: visão negativa de si, do mundo e do futuro.</p>
        <p><strong>Pensamentos automáticos</strong> → emoções → comportamentos (modelo ABC).</p>
        <p>Técnicas: reestruturação cognitiva, registro de pensamentos, exposição, ativação comportamental.</p>
        <p>Beck (2022): Terapia cognitivo-comportamental: teoria e prática.</p>`
      },
      {
        title: "Ellis e REBT",
        content: `<p><strong>Albert Ellis</strong> — Terapia Racional Emotiva (REBT).</p>
        <p>Acontecimento (A) → Crença (B) → Consequência emocional (C).</p>
        <p>Crenças irracionais geram perturbação; disputa racional promove mudança.</p>`
      },
      {
        title: "Legado e Evidências",
        content: `<p>TCC é uma das abordagens com <strong>maior evidência empírica</strong> para diversos transtornos.</p>
        <p>Protocolos manuais, outcome research, integração com neurociências cognitivas.</p>
        <div class="highlight">Críticas: reducionismo, foco em sintoma, pouca profundidade existencial.</div>`
      }
    ],
    quiz: [
      { q: "Beck é principal referência da:", options: ["Terapia cognitivo-comportamental","Psicanálise kleiniana","Gestalt clássica","Psicologia sócio-histórica"], answer: 0, explain: "Beck desenvolveu modelo cognitivo da depressão e TCC." },
      { q: "Na TCC, pensamentos automáticos influenciam:", options: ["Emoções e comportamentos","Apenas estruturas cerebrais","Exclusivamente o inconsciente","Somente reflexos condicionados"], answer: 0, explain: "Modelo cognitivo: cognição medeia afeto e ação." },
      { q: "Ellis criou a:", options: ["REBT — terapia racional emotiva","Teoria de campo de Lewin","Analítica existencial","Psicologia do ego"], answer: 0, explain: "REBT enfatiza disputa de crenças irracionais." },
      { q: "TCC integra principalmente:", options: ["Cognição e comportamento","Fenomenologia e existência","Pulsão e recalque","Representação social"], answer: 0, explain: "Síntese cognitivista e behaviorista com base empírica." },
      { q: "Ativação comportamental é técnica usada em:", options: ["TCC da depressão","Psicanálise freudiana","Hipnose ericksoniana","Análise transacional"], answer: 0, explain: "Aumentar atividades reforçadoras para romper ciclo depressivo." }
    ],
    matching: [
      { term: "Beck", match: "Distorções cognitivas" },
      { term: "Ellis", match: "REBT" },
      { term: "Pensamento automático", match: "Cognição mediadora" },
      { term: "Exposição", match: "Técnica comportamental" },
      { term: "TCC", match: "Evidência empírica" }
    ],
    boss: {
      name: "Arquiteto Cognitivo",
      desc: "Prove domínio da abordagem TCC!",
      questions: [
        { q: "Triade cognitiva de Beck refere-se a:", options: ["Si, mundo e futuro negativos","Id, ego e superego","Percepção, memória e atenção","Estímulo, resposta e reforço"], answer: 0, explain: "Padrão cognitivo central na depressão." },
        { q: "Reestruturação cognitiva busca:", options: ["Modificar crenças disfuncionais","Recalcar pulsões sexuais","Induzir insight gestáltico","Analisar campo social"], answer: 0, explain: "Identificar e questionar pensamentos distorcidos." },
        { q: "TCC surgiu historicamente como:", options: ["Alternativa ao behaviorismo sem cognição","Extensão da psicanálise ortodoxa","Ramo da fenomenologia pura","Derivado do marxismo soviético"], answer: 0, explain: "Cognitivismo respondeu aos limites do behaviorismo radical." },
        { q: "Modelo ABC de Ellis: A significa:", options: ["Acontecimento ativador","Ansiedade automática","Associação behaviorista","Análise existencial"], answer: 0, explain: "A = evento; B = crença; C = consequência emocional." },
        { q: "Crítica frequente à TCC:", options: ["Foco em sintoma e pouca profundidade","Excesso de associação livre","Negar método científico","Rejeitar qualquer experimento"], answer: 0, explain: "Alguns criticam reducionismo e superficialidade existencial." }
      ]
    }
  },
  {
    id: 6,
    title: "Humanismo Americano",
    icon: "🌻",
    color: "#eab308",
    description: "Terceira força: Rogers, Maslow e potencial humano",
    lessons: [
      {
        title: "A Terceira Força",
        content: `<p>O <strong>humanismo</strong> surge nos anos 1950-60 como "terceira força" frente à psicanálise e ao behaviorismo.</p>
        <div class="highlight">Crítica: psicanálise foca patologia; behaviorismo ignora subjetividade. Humanismo valoriza experiência, liberdade e potencial.</div>`
      },
      {
        title: "Carl Rogers",
        content: `<p><strong>Carl Rogers</strong> desenvolveu a <strong>Abordagem Centrada na Pessoa</strong>.</p>
        <ul>
          <li><strong>Condicões de base:</strong> congruência, empatia, consideração positiva incondicional</li>
          <li><strong>Self:</strong> concepção que a pessoa tem de si</li>
          <li><strong>Tendência atualizante:</strong> impulso ao crescimento</li>
        </ul>`
      },
      {
        title: "Abraham Maslow",
        content: `<p><strong>Maslow</strong> propôs a <strong>hierarquia de necessidades</strong> e psicologia da motivação.</p>
        <p>Pirâmide: fisiológicas → segurança → pertencimento → estima → autorrealização.</p>
        <p>Foco em experiências de pico e potencial humano saudável, não só patologia.</p>`
      },
      {
        title: "Movimento Humanista",
        content: `<p>Associação de Psicologia Humanista (AHPA), conferências, influência em educação e organizações.</p>
        <p>Ênfase em método fenomenológico, escuta, autenticidade e ética do encontro.</p>
        <p>Diálogo com existencialismo europeu, mas tom mais otimista e americano.</p>`
      },
      {
        title: "Críticas e Legado",
        content: `<p>Críticas: idealismo, conceitos difíceis de operacionalizar, individualismo cultural.</p>
        <p>Legado: práticas de acolhimento, escuta ativa, coaching, educação humanista.</p>
        <div class="highlight">Influenciou políticas de saúde mental comunitária e direitos do paciente.</div>`
      }
    ],
    quiz: [
      { q: "Humanismo é chamado de terceira força por se opor a:", options: ["Psicanálise e behaviorismo","Gestalt e existencialismo","Marxismo e fenomenologia","TCC e neurociência"], answer: 0, explain: "Alternativa às duas grandes correntes do século XX." },
      { q: "Rogers enfatizou principalmente:", options: ["Empatia e consideração positiva","Reforço e punição","Interpretação de sonhos","Análise de contingências"], answer: 0, explain: "Condicões de base para relação terapêutica." },
      { q: "Maslow é conhecido pela:", options: ["Hierarquia de necessidades","Teoria das relações objetais","Condicionamento operante","Analítica do Dasein"], answer: 0, explain: "Da fisiologia à autorrealização." },
      { q: "Tendência atualizante, para Rogers, é:", options: ["Impulso ao crescimento e desenvolvimento","Recalque de pulsões primitivas","Resposta reflexa condicionada","Determinação genética fixa"], answer: 0, explain: "Força inata em direção à maturidade e autonomia." },
      { q: "Abordagem Centrada na Pessoa prioriza:", options: ["Experiência subjetiva do cliente","Controle experimental rigoroso","Análise de inconsciente coletivo","Mensuração de reflexos"], answer: 0, explain: "Cliente como centro do processo terapêutico." }
    ],
    matching: [
      { term: "Rogers", match: "Centrada na pessoa" },
      { term: "Maslow", match: "Autorrealização" },
      { term: "Terceira força", match: "Humanismo" },
      { term: "Empatia", match: "Condição de base" },
      { term: "Self", match: "Concepção de si" }
    ],
    boss: {
      name: "Guardião Humanista",
      desc: "Demonstre domínio do humanismo americano!",
      questions: [
        { q: "Consideração positiva incondicional significa:", options: ["Aceitação sem julgamento do cliente","Reforço material por comportamento","Interpretação de atos falhos","Análise de estruturas sociais"], answer: 0, explain: "Aceitar o cliente como pessoa, separando avaliação." },
        { q: "Topo da pirâmide de Maslow:", options: ["Autorrealização","Segurança fisiológica","Pertencimento social","Estima básica"], answer: 0, explain: "Nível mais alto: realização do potencial." },
        { q: "Humanismo valoriza especialmente:", options: ["Liberdade e potencial humano","Determinismo biológico total","Patologia e neurose","Controle ambiental estrito"], answer: 0, explain: "Visão otimista do ser humano e sua capacidade de escolha." },
        { q: "Crítica ao humanismo inclui:", options: ["Dificuldade de mensuração científica","Excesso de behaviorismo","Negar subjetividade","Rejeitar clínica"], answer: 0, explain: "Conceitos como self são difíceis de operacionalizar." },
        { q: "Rogers influenciou práticas de:", options: ["Escuta ativa e acolhimento","Eletrochoque e insulinoterapia","Análise de sonhos freudiana","Condicionamento aversivo"], answer: 0, explain: "Escuta empática é marca das práticas humanistas." }
      ]
    }
  },
  {
    id: 7,
    title: "Existencialismo Europeu",
    icon: "🌑",
    color: "#64748b",
    description: "Husserl, Heidegger, Sartre e fenomenologia",
    lessons: [
      {
        title: "Fenomenologia de Husserl",
        content: `<p><strong>Edmund Husserl</strong> funda a fenomenologia: estudo das estruturas da experiência vivida.</p>
        <div class="highlight"><strong>Epoché:</strong> suspender julgamentos naturais para descrever fenômenos como aparecem à consciência.</div>
        <p>Intencionalidade: consciência é sempre consciência de algo.</p>`
      },
      {
        title: "Heidegger e o Dasein",
        content: `<p><strong>Martin Heidegger</strong> desenvolve a <strong>analítica existencial</strong> em Ser e Tempo (1927).</p>
        <p><strong>Dasein</strong> = ser-no-mundo, existência humana lançada em situação.</p>
        <p>Conceitos: ser-para-a-morte, autenticidade, angústia, cuidado (Sorge).</p>`
      },
      {
        title: "Sartre e o Existencialismo",
        content: `<p><strong>Jean-Paul Sartre</strong>: existencialismo francês, ênfase em liberdade e responsabilidade.</p>
        <p>"A existência precede a essência" — não há natureza humana fixa predeterminada.</p>
        <p>Má-fé (mauvaise foi): negar a liberdade e escolher como coisa.</p>`
      },
      {
        title: "Psicologia Fenomenológico-Existencial",
        content: `<p>Influência em psicoterapia existencial (May, Yalom, Frankl).</p>
        <p>Temas: ansiedade existencial, sentido, solidão, liberdade, finitude.</p>
        <p><strong>Logoterapia</strong> (Frankl): busca de sentido como motivação central.</p>`
      },
      {
        title: "Diferenças com Humanismo",
        content: `<p>Existencialismo europeu é mais sombrio: angústia, absurdo, finitude.</p>
        <p>Humanismo americano tende ao otimismo e potencial de crescimento.</p>
        <div class="highlight">Ambos valorizam subjetividade, mas existencialismo enfatiza limites e paradoxos da existência.</div>`
      }
    ],
    quiz: [
      { q: "Husserl é fundador da:", options: ["Fenomenologia","Psicanálise","Behaviorismo radical","Gestalt terapia"], answer: 0, explain: "Fenomenologia descreve estruturas da experiência." },
      { q: "Dasein, para Heidegger, designa:", options: ["Ser humano existente no mundo","Complexo de Édipo infantil","Reflexo condicionado","Representação social"], answer: 0, explain: "Existência humana como ser-para-a-morte no mundo." },
      { q: "Sartre afirmou que a existência:", options: ["Precede a essência","É determinada geneticamente","É puramente inconsciente","Segue leis de reforço"], answer: 0, explain: "Homem é condenado a ser livre — cria sua essência." },
      { q: "Epoché husserliana consiste em:", options: ["Suspender pressupostos naturais","Reforçar comportamentos desejados","Interpretar sonhos simbólicos","Medir tempo de reação"], answer: 0, explain: "Redução fenomenológica para descrever o vivido." },
      { q: "Logoterapia foi desenvolvida por:", options: ["Viktor Frankl","B. F. Skinner","Sigmund Freud","Abraham Maslow"], answer: 0, explain: "Frankl enfatiza busca de sentido mesmo no sofrimento." }
    ],
    matching: [
      { term: "Husserl", match: "Fenomenologia" },
      { term: "Heidegger", match: "Dasein" },
      { term: "Sartre", match: "Existência precede essência" },
      { term: "Frankl", match: "Logoterapia" },
      { term: "Angústia", match: "Confronto com liberdade" }
    ],
    boss: {
      name: "Sombras do Ser",
      desc: "Enfrente o chefe existencial!",
      questions: [
        { q: "Ser-para-a-morte (Heidegger) refere-se a:", options: ["Consciência da finitude existencial","Reflexo medular espinal","Complexo castrador","Reforço intermitente"], answer: 0, explain: "Antecipação da morte estrutura existência autêntica." },
        { q: "Má-fé sartreana é:", options: ["Negar a própria liberdade","Recalcar memórias traumáticas","Organizar percepção em gestalt","Seguir hierarquia de Maslow"], answer: 0, explain: "Fingir ser coisa determinada, fugindo da liberdade." },
        { q: "Intencionalidade significa que consciência:", options: ["É sempre de algo","É tabula rasa vazia","É puramente reflexa","Não existe fenomenologicamente"], answer: 0, explain: "Consciência aponta para objetos de experiência." },
        { q: "Yalom trabalha temas como:", options: ["Morte, liberdade, isolamento e sentido","Reforço, punição e extinção","Pulsão, recalque e sublimação","Campo, insight e percepção"], answer: 0, explain: "Preocupações existenciais na psicoterapia." },
        { q: "Existencialismo europeu contrasta com humanismo por:", options: ["Maior ênfase em angústia e finitude","Rejeição total de subjetividade","Adesão ao behaviorismo","Foco em reforço positivo"], answer: 0, explain: "Tom mais trágico e filosófico que o otimismo humanista." }
      ]
    }
  },
  {
    id: 8,
    title: "Psicologia Sócio-Histórica",
    icon: "📖",
    color: "#ef4444",
    description: "Marx, Vygotsky e diversidade das psicologias",
    lessons: [
      {
        title: "Contribuição Marxista",
        content: `<p><strong>Materialismo Histórico e Dialético</strong>: consciência é produzida nas relações sociais de produção.</p>
        <p>Psique não é dado natural isolado — é historicamente constituída.</p>
        <div class="highlight">Marx: "Os homens fazem sua história, mas não a fazem como querem" — determinação social e agency.</div>`
      },
      {
        title: "Primórdios Sócio-Históricos",
        content: `<p>Na URSS pós-revolução, psicologia busca base materialista e social.</p>
        <p><strong>Vygotsky</strong> (1896-1934): zona de desenvolvimento proximal, mediação cultural, linguagem.</p>
        <p>Crítica a métodos introspectivos e isolamento do contexto social.</p>`
      },
      {
        title: "Vygotsky e Mediação",
        content: `<p>Desenvolvimento psíquico superior ocorre por <strong>mediação</strong> de signos e instrumentos culturais.</p>
        <ul>
          <li><strong>ZDP:</strong> distância entre desempenho atual e potencial com auxílio</li>
          <li>Funções psíquicas inferiores → superiores (voluntárias)</li>
          <li>Pensamento e linguagem: internalização de formas sociais</li>
        </ul>`
      },
      {
        title: "Concretização no Brasil",
        content: `<p>Psicologia sócio-histórica no Brasil: Gonzalez Rey, Mitjáns, Colinvaux, entre outros.</p>
        <p>Diálogo com educação, saúde coletiva e políticas públicas.</p>
        <p>Ênfase em subjetividade como produção histórica, não interioridade privada.</p>`
      },
      {
        title: "Psicologia ou Psicologias?",
        content: `<p>Encerramento: reconhecer <strong>pluralidade teórica</strong> sem relativismo ingênuo.</p>
        <p>Cada matriz oferece ferramentas e limites; prática exige escolha reflexiva e ética.</p>
        <div class="highlight">Psicologia é saber referido às condições históricas, filosóficas e sociais de sua produção.</div>`
      }
    ],
    quiz: [
      { q: "Vygotsky é referência central da:", options: ["Psicologia sócio-histórica","Behaviorismo radical","Psicanálise kleiniana","Gestalt de percepção"], answer: 0, explain: "Mediação cultural e desenvolvimento na ZDP." },
      { q: "Zona de Desenvolvimento Proximal é:", options: ["Distância entre atual e potencial com ajuda","Período de latência freudiana","Estágio de autorrealização","Fase de negação existencial"], answer: 0, explain: "O que a criança faz com mediação de adulto ou par." },
      { q: "Materialismo histórico enfatiza:", options: ["Condições sociais na constituição da psique","Determinismo biológico puro","Inconsciente como única causa","Reforço como explicação total"], answer: 0, explain: "Subjetividade é produzida historicamente nas relações sociais." },
      { q: "Para Vygotsky, linguagem e pensamento:", options: ["Desenvolvem-se em relação dialética","São totalmente independentes","São reflexos inatos fixos","Não importam ao desenvolvimento"], answer: 0, explain: "Linguagem social internaliza-se como pensamento." },
      { q: "Tema 'Psicologia ou Psicologias?' enfatiza:", options: ["Pluralidade teórica historicamente situada","Unicidade de método experimental","Rejeição de toda clínica","Identidade com neurologia"], answer: 0, explain: "Reconhecer diversidade sem apagar diferenças epistemológicas." }
    ],
    matching: [
      { term: "Vygotsky", match: "Zona de desenvolvimento proximal" },
      { term: "Marx", match: "Materialismo histórico" },
      { term: "Mediação cultural", match: "Signos e instrumentos" },
      { term: "Sócio-histórica", match: "Subjetividade social" },
      { term: "Internalização", match: "Formas sociais no indivíduo" }
    ],
    boss: {
      name: "Guardião Final",
      desc: "O chefe final! Prove que domina todo o semestre!",
      questions: [
        { q: "Vygotsky criticou métodos que:", options: ["Isolam indivíduo do contexto social","Usam qualquer observação","Valorizam mediação cultural","Estudam linguagem na infância"], answer: 0, explain: "Psique não se compreende fora das relações sociais históricas." },
        { q: "Funções psíquicas superiores, para Vygotsky:", options: ["São mediadas culturalmente","São puramente reflexas","Existem ao nascer completas","Não dependem de linguagem"], answer: 0, explain: "Voluntárias, arbitrárias, internalizadas da cultura." },
        { q: "Psicologia sócio-histórica no Brasil articula-se com:", options: ["Educação e políticas públicas","Neurocirurgia estereotáxica","Farmacologia psiquiátrica apenas","Estatística matemática pura"], answer: 0, explain: "Aplicações em formação, saúde e intervenção social." },
        { q: "Dialética marxista implica compreender:", options: ["Contradições e transformação histórica","Estímulo e resposta fixos","Inconsciente imutável","Percepção como único dado"], answer: 0, explain: "Realidade em movimento, contradição e totalidade." },
        { q: "Encerrar o curso significa reconhecer que:", options: ["Cada teoria tem contexto, limites e contribuições","Uma abordagem explica tudo","História não importa à ciência","Método é idêntico em todas as escolas"], answer: 0, explain: "Competência: comparar matrizes com criticidade histórica." }
      ]
    }
  }
];

const ACHIEVEMENTS = [
  { id: "first_lesson", icon: "📖", title: "Primeira Lição", desc: "Complete sua primeira lição", xp: 0 },
  { id: "first_quiz", icon: "✅", title: "Quizzer", desc: "Complete seu primeiro quiz", xp: 0 },
  { id: "first_boss", icon: "⚔️", title: "Caçador de Chefes", desc: "Derrote seu primeiro chefe", xp: 0 },
  { id: "unit1", icon: "🏛️", title: "Historiador", desc: "Complete a Unidade 1", xp: 0 },
  { id: "unit4", icon: "🛋️", title: "Analista", desc: "Complete a Unidade 4", xp: 0 },
  { id: "unit8", icon: "🏆", title: "Mestre das Matrizes", desc: "Complete todas as 8 unidades", xp: 0 },
  { id: "streak5", icon: "🔥", title: "Em Chamas", desc: "Acerte 5 questões seguidas", xp: 0 },
  { id: "streak10", icon: "💎", title: "Imparável", desc: "Acerte 10 questões seguidas", xp: 0 },
  { id: "no_life_lost", icon: "🛡️", title: "Invencível", desc: "Complete um quiz sem perder vidas", xp: 0 },
  { id: "level5", icon: "⭐", title: "Veterano", desc: "Alcance o nível 5", xp: 0 },
  { id: "level10", icon: "👑", title: "Lenda Teórica", desc: "Alcance o nível 10", xp: 0 },
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

const UNITS = [
  {
    id: 1,
    title: "Ciência Psicológica",
    icon: "🔬",
    color: "#ec4899",
    description: "Senso comum, método científico e investigação",
    lessons: [
      {
        title: "Senso Comum vs Ciência",
        content: `<p>A <strong>Psicologia científica</strong> diferencia-se do senso comum por usar métodos sistemáticos, controle e revisão por pares.</p>
        <div class="highlight">Senso comum = crenças populares. Ciência = hipóteses testáveis, evidência e replicabilidade.</div>
        <p>Feldman e Myers são referências básicas da disciplina.</p>`
      },
      {
        title: "Método Científico",
        content: `<p>Etapas: <strong>observação → hipótese → experimento → análise → conclusão → publicação</strong>.</p>
        <ul>
          <li><strong>Variáveis:</strong> independente (manipulada) e dependente (medida)</li>
          <li><strong>Controle:</strong> isolar fatores confundidores</li>
          <li><strong>Replicação:</strong> validar resultados</li>
        </ul>`
      },
      {
        title: "Métodos de Investigação",
        content: `<p><strong>Experimental:</strong> causa e efeito, laboratório ou campo.</p>
        <p><strong>Correlacional:</strong> relações entre variáveis, sem inferir causalidade.</p>
        <p><strong>Descritivo:</strong> estudos de caso, observação, entrevistas.</p>
        <p>Competência: planejar coleta de dados e usar ferramentas informatizadas.</p>`
      },
      {
        title: "Ética em Pesquisa",
        content: `<p><strong>TCLE</strong> (Termo de Consentimento Livre e Esclarecido) obrigatório em pesquisas com humanos.</p>
        <p>No TP: não usar a palavra "teste"; tranquilizar participantes; não há certo/errado.</p>
        <div class="highlight">Vulneráveis excluídos do TP: crianças, adolescentes, pessoas em vulnerabilidade ou com transtorno mental.</div>`
      },
      {
        title: "Trabalho Prático (TP)",
        content: `<p>TP em 4 etapas: tema (Sensação, Percepção ou Memória), projeto, coleta (20 participantes) e apresentação oral.</p>
        <p>Dois grupos comparativos (ex.: jovens x idosos). Instrumentos simples: listas de palavras, figuras ambíguas.</p>
        <p>Relatório ~15 páginas em ABNT: Resumo, Introdução, Método, Resultados, Discussão, Conclusão.</p>`
      }
    ],
    quiz: [
      { q: "A ciência psicológica diferencia-se do senso comum por:", options: ["Usar métodos sistemáticos e evidência","Aceitar crenças sem questionar","Ignorar dados empíricos","Rejeitar toda observação"], answer: 0, explain: "Ciência exige método, controle e evidência testável." },
      { q: "Variável independente é:", options: ["Manipulada pelo pesquisador","Sempre o resultado medido","Impossível de controlar","Igual à dependente"], answer: 0, explain: "VI é a causa ou condição manipulada no experimento." },
      { q: "TCLE significa:", options: ["Termo de Consentimento Livre e Esclarecido","Teste Cognitivo de Leitura Escrita","Técnica Clínica de Liberação Emocional","Tratamento Comportamental de Longa Evolução"], answer: 0, explain: "Documento ético obrigatório antes da coleta de dados." },
      { q: "No TP, temas aceitos são:", options: ["Sensação, Percepção ou Memória","Qualquer tema livre","Apenas Psicanálise","Somente Motivação"], answer: 0, explain: "Plano de ensino restringe temas aos processos estudados." },
      { q: "Método correlacional estabelece:", options: ["Relações entre variáveis sem causalidade direta","Causa com certeza absoluta","Diagnóstico clínico formal","Prescrição de medicamentos"], answer: 0, explain: "Correlação não implica causalidade." }
    ],
    matching: [
      { term: "Hipótese", match: "Previsão testável" },
      { term: "TCLE", match: "Consentimento informado" },
      { term: "Variável dependente", match: "Resultado medido" },
      { term: "Replicação", match: "Repetir estudo" },
      { term: "Senso comum", match: "Crença popular" }
    ],
    boss: {
      name: "Guardião Científico",
      desc: "Prove domínio dos fundamentos da ciência psicológica!",
      questions: [
        { q: "Cada grupo no TP deve coletar dados de:", options: ["20 participantes no total","5 participantes apenas","100 participantes obrigatórios","Nenhum participante humano"], answer: 0, explain: "10 em cada grupo de comparação." },
        { q: "Etapa 2 do TP inclui:", options: ["Introdução, Método e Referências","Apenas apresentação oral","Coleta sem aprovação prévia","Resultados finais completos"], answer: 0, explain: "Etapa 2 é teórica — projeto sem coleta ainda." },
        { q: "Feldman é autor de:", options: ["Introdução à Psicologia","Representações Sociais","Matrizes do Pensamento","Neuroanatomia Clínica"], answer: 0, explain: "Obra básica da bibliografia." },
        { q: "Ao contatar participantes, evita-se:", options: ["Usar a palavra teste","Explicar o objetivo da pesquisa","Pedir consentimento","Registrar dados objetivamente"], answer: 0, explain: "Orientação ética do plano de ensino." },
        { q: "Competência da disciplina inclui:", options: ["Identificar processos em exemplos cotidianos","Realizar cirurgia neurológica","Prescrever psicofármacos","Ignorar método científico"], answer: 0, explain: "Aplicar conceitos a situações do dia a dia." }
      ]
    }
  },
  {
    id: 2,
    title: "Sensação",
    icon: "👁️",
    color: "#f43f5e",
    description: "Sentidos, limiares, transdução e adaptação",
    lessons: [
      {
        title: "O Processo Sensorial",
        content: `<p><strong>Sensação</strong> = processo pelo qual estímulos físicos são captados e convertidos em sinais neurais.</p>
        <p><strong>Neurônios sensoriais</strong> especializados em cada modalidade (visão, audição, etc.).</p>
        <div class="highlight"><strong>Transdução:</strong> energia física → impulso elétrico neural.</div>`
      },
      {
        title: "Limiares Sensoriais",
        content: `<p><strong>Limiar absoluto:</strong> intensidade mínima para detectar estímulo 50% das vezes.</p>
        <p><strong>Limiar diferencial (DL):</strong> menor diferença perceptível entre dois estímulos (Lei de Weber).</p>
        <p>Sensibilidade varia com atenção, fadiga e contexto.</p>`
      },
      {
        title: "Adaptação Sensorial",
        content: `<p><strong>Adaptação:</strong> diminuição da resposta sensorial a estímulo constante.</p>
        <p>Exemplo: não perceber o relógio após alguns minutos; cheiro que "some".</p>
        <p>Mecanismo de economia atencional — prioriza mudanças no ambiente.</p>`
      },
      {
        title: "Os Cinco Sentidos",
        content: `<ul>
          <li><strong>Visão:</strong> retina, cones e bastonetes, cores</li>
          <li><strong>Audição:</strong> cóclea, frequência e timbre</li>
          <li><strong>Tato:</strong> pressão, temperatura, dor</li>
          <li><strong>Olfato:</strong> receptores químicos no nariz</li>
          <li><strong>Gustação:</strong> papilas gustativas (doce, salgado, azedo, amargo, umami)</li>
        </ul>`
      },
      {
        title: "Sentidos Complementares",
        content: `<p><strong>Propriocepção:</strong> posição do corpo no espaço.</p>
        <p><strong>Vestibular:</strong> equilíbrio e aceleração (ouvido interno).</p>
        <p><strong>Cinestesia:</strong> movimento dos músculos e articulações.</p>
        <div class="highlight">Sensação é entrada de dados; percepção organiza e interpreta.</div>`
      }
    ],
    quiz: [
      { q: "Transdução sensorial é:", options: ["Conversão de estímulo em sinal neural","Esquecimento de informações","Organização gestáltica de formas","Motivação por reforço"], answer: 0, explain: "Receptores convertem energia física em impulso nervoso." },
      { q: "Limiar absoluto indica:", options: ["Intensidade mínima detectável","Diferença entre dois estímulos","Tempo de reação motor","Capacidade de memória"], answer: 0, explain: "Ponto em que estímulo é detectado metade das vezes." },
      { q: "Adaptação sensorial ocorre quando:", options: ["Resposta diminui a estímulo constante","Memória de longo prazo falha","Emoção intensifica percepção","Motivação aumenta fome"], answer: 0, explain: "Sistema sensorial se ajusta a estímulos invariáveis." },
      { q: "Cones na retina são responsáveis por:", options: ["Visão de cores e detalhes","Audição de altas frequências","Equilíbrio vestibular","Olfato químico"], answer: 0, explain: "Cones = visão diurna e cor; bastonetes = pouca luz." },
      { q: "Limiar diferencial relaciona-se à:", options: ["Lei de Weber","Pirâmide de Maslow","Teoria de James-Lange","Heurística de disponibilidade"], answer: 0, explain: "Weber: ΔI/I = constante para detectar diferença." }
    ],
    matching: [
      { term: "Transdução", match: "Estímulo → neural" },
      { term: "Limiar absoluto", match: "Detecção mínima" },
      { term: "Adaptação", match: "Resposta diminui" },
      { term: "Cóclea", match: "Audição" },
      { term: "Propriocepção", match: "Posição corporal" }
    ],
    boss: {
      name: "Sensor Supremo",
      desc: "Enfrente o chefe da sensação!",
      questions: [
        { q: "Bastonetes na retina funcionam melhor em:", options: ["Pouca luminosidade","Visão de cores finas","Audição musical","Olfato agudo"], answer: 0, explain: "Bastonetes = visão escotópica (noturna)." },
        { q: "Sentido vestibular está no:", options: ["Ouvido interno","Córtex occipital","Hipocampo medial","Cerebelo apenas"], answer: 0, explain: "Sistema vestibular regula equilíbrio." },
        { q: "Sensação difere de percepção porque:", options: ["Sensação é registro bruto do estímulo","São processos idênticos","Percepção não usa sentidos","Sensação é sempre ilusória"], answer: 0, explain: "Sensação = entrada; percepção = organização e significado." },
        { q: "Papilas gustativas detectam:", options: ["Sabores básicos","Cores visuais","Sons graves","Movimento profundo"], answer: 0, explain: "Doce, salgado, azedo, amargo, umami." },
        { q: "Testes sensoriais no TP podem incluir:", options: ["Avaliação dos sentidos","Cirurgia experimental","EEG de sono profundo","Tomografia computadorizada"], answer: 0, explain: "Instrumentos simples ligados a sensação/percepção." }
      ]
    }
  },
  {
    id: 3,
    title: "Percepção",
    icon: "🎨",
    color: "#8b5cf6",
    description: "Gestalt, atenção, profundidade e processamento",
    lessons: [
      {
        title: "Definição e Organização",
        content: `<p><strong>Percepção</strong> = organização e interpretação de sensações em experiência significativa.</p>
        <p>Percepção de formas, padrões, objetos e cores vai além dos dados sensoriais brutos.</p>
        <div class="highlight">O cérebro preenche lacunas e organiza o mundo em figuras e fundos.</div>`
      },
      {
        title: "Leis da Gestalt",
        content: `<ul>
          <li><strong>Proximidade:</strong> elementos próximos agrupam-se</li>
          <li><strong>Similaridade:</strong> elementos semelhantes formam grupo</li>
          <li><strong>Continuidade:</strong> preferência por linhas suaves</li>
          <li><strong>Fechamento:</strong> completar formas incompletas</li>
        </ul>
        <p>Percepção é ativa, não passiva.</p>`
      },
      {
        title: "Profundidade e Movimento",
        content: `<p><strong>Pistas monocular:</strong> sobreposição, perspectiva linear, textura gradiente.</p>
        <p><strong>Pistas binoculares:</strong> convergência e disparidade retiniana.</p>
        <p><strong>Percepção de movimento:</strong> phi phenomenon, movimento induzido.</p>
        <p><strong>Ilusões:</strong> óticas e de movimento demonstram construção perceptiva.</p>`
      },
      {
        title: "Constância e Predisposição",
        content: `<p><strong>Constância perceptiva:</strong> tamanho, forma e cor parecem estáveis apesar de mudanças na imagem retinal.</p>
        <p><strong>Predisposição:</strong> expectativas e contexto influenciam o que percebemos.</p>
        <p>Exemplo: figuras ambíguas (jovem/velha) — uso no TP.</p>`
      },
      {
        title: "Atenção e Processamento",
        content: `<p><strong>Atenção concentrada:</strong> foco em um estímulo.</p>
        <p><strong>Alternada:</strong> alternar entre tarefas/estímulos.</p>
        <p><strong>Sustentada:</strong> manter foco por período prolongado.</p>
        <p><strong>Bottom-up:</strong> dados sensoriais dirigem percepção.</p>
        <p><strong>Top-down:</strong> expectativas e conhecimento dirigem percepção.</p>`
      }
    ],
    quiz: [
      { q: "Lei da proximidade da Gestalt diz que:", options: ["Elementos próximos agrupam-se","Cores opostas se atraem","Movimento sempre é real","Memória é infinita"], answer: 0, explain: "Proximidade espacial cria agrupamento perceptivo." },
      { q: "Processamento top-down significa:", options: ["Expectativas guiam percepção","Apenas dados sensoriais brutos","Ausência total de cognição","Reflexo medular puro"], answer: 0, explain: "Conhecimento prévio influencia interpretação." },
      { q: "Constância de tamanho faz com que:", options: ["Objetos pareçam do mesmo tamanho à distância","Tudo pareça maior de noite","Cores mudem aleatoriamente","Sons fiquem mais agudos"], answer: 0, explain: "Percebemos tamanho real apesar da imagem retinal menor." },
      { q: "Atenção sustentada é:", options: ["Manter foco por longo período","Alternar entre tarefas rapidamente","Dormir profundamente","Esquecer estímulos novos"], answer: 0, explain: "Vigilância contínua — exige esforço cognitivo." },
      { q: "Ilusões de ótica demonstram que:", options: ["Percepção é construção ativa","Sentidos são perfeitos","Memória não existe","Motivação é instinto"], answer: 0, explain: "Cérebro interpreta — às vezes erroneamente." }
    ],
    matching: [
      { term: "Gestalt", match: "Leis de organização" },
      { term: "Top-down", match: "Expectativas guiam" },
      { term: "Bottom-up", match: "Dados sensoriais" },
      { term: "Constância", match: "Estabilidade perceptiva" },
      { term: "Figura ambígua", match: "Dupla interpretação" }
    ],
    boss: {
      name: "Ilusionista Mestre",
      desc: "Demonstre domínio da percepção!",
      questions: [
        { q: "Disparidade retiniana ajuda na percepção de:", options: ["Profundidade binocular","Audição estéreo","Olfato químico","Memória episódica"], answer: 0, explain: "Dois olhos — imagens ligeiramente diferentes — profundidade." },
        { q: "Fechamento gestáltico consiste em:", options: ["Completar formas incompletas","Fechar os olhos para dormir","Encerrar experimento","Recalcar memórias"], answer: 0, explain: "Tendência a ver figuras fechadas." },
        { q: "Figuras ambíguas no TP avaliam:", options: ["Percepção","Motivação sexual","Sono REM","Linguagem articulada"], answer: 0, explain: "Tema de pesquisa permitido: Percepção." },
        { q: "Phi phenomenon relaciona-se a:", options: ["Percepção de movimento","Fome e saciedade","Emoção básica","Pensamento lógico"], answer: 0, explain: "Ilusão de movimento com luzes sequenciais." },
        { q: "Atenção alternada envolve:", options: ["Trocar foco entre estímulos","Manter foco sem pausa","Dormir e sonhar","Esquecer deliberadamente"], answer: 0, explain: "Multitarefa atencional — custo cognitivo." }
      ]
    }
  },
  {
    id: 4,
    title: "Memória",
    icon: "💾",
    color: "#3b82f6",
    description: "Codificação, tipos de memória e esquecimento",
    lessons: [
      {
        title: "Modelo de Processamento",
        content: `<p>Três estágios: <strong>codificação → armazenamento → recuperação</strong>.</p>
        <p><strong>Memória sensorial:</strong> registro brevíssimo (ícones, ecos).</p>
        <p><strong>Memória de curto prazo (MCp):</strong> ~20-30 seg, capacidade limitada (7±2 itens).</p>
        <p><strong>Memória de longo prazo (MLp):</strong> armazenamento duradouro.</p>`
      },
      {
        title: "Tipos de Memória",
        content: `<ul>
          <li><strong>Explícita (declarativa):</strong> episódica (eventos) e semântica (fatos)</li>
          <li><strong>Implícita (não declarativa):</strong> procedural (habilidades), priming</li>
          <li><strong>Memória flash:</strong> vívidas de eventos emocionais (11/9, etc.)</li>
        </ul>
        <p>Sacks (O homem que confundiu sua mulher com um chapéu) ilustra distúrbios de memória.</p>`
      },
      {
        title: "Estratégias de Codificação",
        content: `<p><strong>Repetição, elaboração, organização, mnemônicos</strong> melhoram codificação.</p>
        <p><strong>Chunking:</strong> agrupar informações (ex.: telefone em blocos).</p>
        <p>Níveis de processamento (Craik & Lockhart): processamento profundo = melhor memória.</p>`
      },
      {
        title: "Esquecimento",
        content: `<p><strong>Decaimento:</strong> informação não usada se deteriora.</p>
        <p><strong>Interferência:</strong> proativa (antiga atrapalha nova) e retroativa (nova atrapalha antiga).</p>
        <p><strong>Falha de recuperação:</strong> informação existe mas não é acessada (ponta da língua).</p>
        <p>Curva do esquecimento (Ebbinghaus): maior perda nas primeiras horas.</p>`
      },
      {
        title: "Memória no TP",
        content: `<p>Listas de palavras são instrumento clássico para avaliar memória no TP.</p>
        <p>Comparar grupos (ex.: quem usa óculos x quem não usa) em desempenho de recordação.</p>
        <div class="highlight">20 participantes, 2 grupos de 10, TCLE assinado, análise de resultados em tabelas/gráficos.</div>`
      }
    ],
    quiz: [
      { q: "Memória de curto prazo tem capacidade aproximada de:", options: ["7±2 itens","Infinitos itens","1 item apenas","1000 itens fixos"], answer: 0, explain: "Miller (1956): magical number seven." },
      { q: "Memória episódica armazena:", options: ["Eventos autobiográficos","Procedimentos motores","Reflexos condicionados","Habilidades de bicicleta"], answer: 0, explain: "Episódica = experiências pessoais com contexto." },
      { q: "Memória procedural é tipo de:", options: ["Memória implícita","Memória semântica apenas","Memória sensorial","Memória flash"], answer: 0, explain: "Habilidades como andar, digitar — sem recordação consciente." },
      { q: "Curva do esquecimento foi proposta por:", options: ["Ebbinghaus","Freud","Skinner","Maslow"], answer: 0, explain: "Ebbinghaus estudou esquecimento de sílabas sem sentido." },
      { q: "Interferência retroativa ocorre quando:", options: ["Nova informação atrapalha antiga","Antiga atrapalha apenas nova","Não há esquecimento","Memória é perfeita"], answer: 0, explain: "Aprender lista B prejudica recordação da lista A." }
    ],
    matching: [
      { term: "Codificação", match: "Registrar informação" },
      { term: "Memória semântica", match: "Fatos e conceitos" },
      { term: "Chunking", match: "Agrupar itens" },
      { term: "Ebbinghaus", match: "Curva do esquecimento" },
      { term: "Memória flash", match: "Evento emocional vívido" }
    ],
    boss: {
      name: "Arquivista Mental",
      desc: "Enfrente o chefe da memória!",
      questions: [
        { q: "Sacks descreve casos de:", options: ["Distúrbios neurológicos e memória","Behaviorismo radical","Representações sociais","Materialismo histórico"], answer: 0, explain: "O homem que confundiu sua mulher com um chapéu." },
        { q: "Memória de trabalho relaciona-se à:", options: ["MCp ativa e manipulação","Sono profundo apenas","Motivação sexual","Percepção de cores"], answer: 0, explain: "Baddeley: sistema executivo central + loop + esboço visuoespacial." },
        { q: "Falha de recuperação (ponta da língua) sugere:", options: ["Informação existe mas não é acessada","Memória foi totalmente apagada","Sensação está ausente","Emoção não existe"], answer: 0, explain: "Problema de acesso, não necessariamente de armazenamento." },
        { q: "Listas de palavras no TP testam principalmente:", options: ["Memória","Motivação","Sono REM","Linguagem escrita apenas"], answer: 0, explain: "Tema permitido: Memória." },
        { q: "Processamento profundo de Craik & Lockhart melhora:", options: ["Codificação e retenção","Apenas velocidade motora","Sono profundo","Percepção de movimento"], answer: 0, explain: "Elaboração semântica fortalece memória." }
      ]
    }
  },
  {
    id: 5,
    title: "Estados de Consciência",
    icon: "😴",
    color: "#6366f1",
    description: "Sono, sonho, ritmo circadiano e alterações",
    lessons: [
      {
        title: "Vigília e Consciência",
        content: `<p><strong>Estado de vigília:</strong> consciência clara, atenção e orientação no tempo/espaço.</p>
        <p>Consciência = awareness subjetivo da experiência e do ambiente.</p>
        <p>Flutua ao longo do dia (alerta → sonolência).</p>`
      },
      {
        title: "Sono e Estágios",
        content: `<p><strong>NREM:</strong> estágios 1-4, sono profundo, restauração física.</p>
        <p><strong>REM:</strong> sonhos vívidos, paralisia muscular, atividade cerebral alta.</p>
        <p>Ciclo ~90 minutos, 4-6 ciclos por noite.</p>
        <div class="highlight">Privação de sono prejudica atenção, memória e humor.</div>`
      },
      {
        title: "Ritmo Circadiano",
        content: `<p><strong>Ritmo circadiano:</strong> ciclo biológico ~24h regulado pelo núcleo supraquiasmático.</p>
        <p>Influenciado por luz (melatonina) e hábitos sociais.</p>
        <p>Jet lag e trabalho noturno desalinham o relógio biológico.</p>`
      },
      {
        title: "Sonho",
        content: `<p>Teorias: processamento emocional (Freud), consolidação de memória, ativação-síntese.</p>
        <p>Sonho REM: narrativas bizaras; NREM: pensamentos mais lógicos.</p>
        <p>Conteúdo influenciado por experiências, emoções e estresse.</p>`
      },
      {
        title: "Estados Alterados",
        content: `<p><strong>Coma:</strong> perda de consciência e resposta ao ambiente.</p>
        <p><strong>Anestesia:</strong> supressão temporária da consciência para cirurgia.</p>
        <p><strong>Substâncias psicoativas:</strong></p>
        <ul>
          <li><strong>Depressoras:</strong> álcool, barbitúricos — reduzem atividade</li>
          <li><strong>Estimulantes:</strong> cafeína, anfetaminas — aumentam alerta</li>
          <li><strong>Alucinógenas:</strong> LSD, psilocibina — alteram percepção</li>
        </ul>`
      }
    ],
    quiz: [
      { q: "Sono REM caracteriza-se por:", options: ["Sonhos vívidos e paralisia muscular","Sono mais profundo e sem sonhos","Vigília total prolongada","Ausência de atividade cerebral"], answer: 0, explain: "REM = Rapid Eye Movement, alta ativação cortical." },
      { q: "Ritmo circadiano dura aproximadamente:", options: ["24 horas","90 minutos","7 dias","1 hora"], answer: 0, explain: "Ciclo dia-noite biológico." },
      { q: "Melatonina está relacionada a:", options: ["Indução do sono","Motivação sexual","Memória procedural","Percepção de cores"], answer: 0, explain: "Hormônio da pineal — escuridão aumenta produção." },
      { q: "Álcool é substância:", options: ["Depressora do SNC","Estimulante pura","Alucinógena clássica","Vitaminas do sono"], answer: 0, explain: "Deprime atividade neural — sedação." },
      { q: "Privação de sono afeta principalmente:", options: ["Atenção e memória","Visão de cores apenas","Audição pura","Motivação instintiva"], answer: 0, explain: "Funções cognitivas e emocionais deterioram." }
    ],
    matching: [
      { term: "REM", match: "Sonhos vívidos" },
      { term: "Circadiano", match: "Ciclo de 24h" },
      { term: "Melatonina", match: "Indução do sono" },
      { term: "Coma", match: "Consciência ausente" },
      { term: "Estimulante", match: "Aumenta alerta" }
    ],
    boss: {
      name: "Senhor do Sono",
      desc: "Demonstre domínio dos estados de consciência!",
      questions: [
        { q: "Núcleo supraquiasmático regula:", options: ["Ritmo circadiano","Audição musical","Memória episódica","Motivação de fome"], answer: 0, explain: "Relógio biológico no hipotálamo." },
        { q: "Anestesia geral produz:", options: ["Perda temporária de consciência","Memória flash permanente","Motivação extrema","Percepção gestáltica"], answer: 0, explain: "Supressão da consciência para procedimentos." },
        { q: "LSD é classificado como:", options: ["Alucinógeno","Depressor","Estimulante leve","Vitaminas"], answer: 0, explain: "Altera percepção, pensamento e emoção." },
        { q: "Ciclo completo do sono dura cerca de:", options: ["90 minutos","24 horas","5 minutos","12 horas"], answer: 0, explain: "NREM + REM repetem-se na noite." },
        { q: "Jet lag desorganiza principalmente:", options: ["Ritmo circadiano","Memória de longo prazo","Percepção de profundidade","Motivação sexual"], answer: 0, explain: "Mudança de fuso horário desalinha relógio biológico." }
      ]
    }
  },
  {
    id: 6,
    title: "Motivação",
    icon: "🔥",
    color: "#f59e0b",
    description: "Teorias, necessidades, fome e excitação",
    lessons: [
      {
        title: "Conceito de Motivação",
        content: `<p><strong>Motivação</strong> = processos que energizam, direcionam e sustentam comportamento em direção a metas.</p>
        <p>Distingue comportamento dirigido a objetivo de reflexos automáticos.</p>
        <p>Pode ser intrínseca (prazer interno) ou extrínseca (recompensa externa).</p>`
      },
      {
        title: "Teorias Clássicas",
        content: `<ul>
          <li><strong>Instintos (McDougal):</strong> comportamento inato — criticada</li>
          <li><strong>Redução do impulso (Hull):</strong> motivação para reduzir tensão/homeostase</li>
          <li><strong>Ótimo de excitação (Hebb):</strong> buscamos nível ideal de arousal</li>
        </ul>`
      },
      {
        title: "Pirâmide de Maslow",
        content: `<p>Do básico ao superior: fisiológicas → segurança → pertencimento → estima → autorrealização.</p>
        <p>Crítica: hierarquia rígida nem sempre se confirma empiricamente.</p>
        <div class="highlight">Maslow enfatiza motivação humanística e potencial de crescimento.</div>`
      },
      {
        title: "Motivação da Fome",
        content: `<p><strong>Homeostase:</strong> hipotálamo regula fome e saciedade (leptina, grelina, glicose).</p>
        <p>Fatores: biológicos, psicológicos (estresse, emoção) e culturais (normas alimentares).</p>
        <p>Comer emocional e distúrbios alimentares envolvem motivação complexa.</p>`
      },
      {
        title: "Motivação Sexual",
        content: `<p>Influenciada por hormônios (testosterona, estrogênio), cognição, cultura e aprendizagem.</p>
        <p>Kinsey, Masters & Johnson contribuíram para estudo científico da sexualidade.</p>
        <p>Motivação sexual não é puramente instintiva — varia entre indivíduos e contextos.</p>`
      }
    ],
    quiz: [
      { q: "Teoria da redução do impulso (Hull) propõe:", options: ["Comportamento reduz tensão interna","Motivação é puramente cognitiva","Instintos explicam tudo","Não há motivação humana"], answer: 0, explain: "Drive reduction: busca homeostase." },
      { q: "Maslow colocou no topo da pirâmide:", options: ["Autorrealização","Fome fisiológica","Segurança básica","Sono profundo"], answer: 0, explain: "Nível mais alto das necessidades humanas." },
      { q: "Ótimo de excitação sugere que buscamos:", options: ["Nível ideal de arousal","Zero estímulo sempre","Máxima apatia","Apenas instintos"], answer: 0, explain: "Hebb: muito pouco ou muito estímulo é desconfortável." },
      { q: "Hipotálamo regula principalmente:", options: ["Fome e saciedade","Visão de cores","Memória episódica","Linguagem articulada"], answer: 0, explain: "Centro de homeostase energética." },
      { q: "Motivação intrínseca vem de:", options: ["Prazer interno na atividade","Recompensa externa apenas","Punição social","Obrigação legal"], answer: 0, explain: "Fazer algo pelo interesse, não por prêmio." }
    ],
    matching: [
      { term: "Maslow", match: "Pirâmide de necessidades" },
      { term: "Hull", match: "Redução do impulso" },
      { term: "Hebb", match: "Ótimo de excitação" },
      { term: "Homeostase", match: "Equilíbrio interno" },
      { term: "Leptina", match: "Sinal de saciedade" }
    ],
    boss: {
      name: "Motor da Ação",
      desc: "Enfrente o chefe da motivação!",
      questions: [
        { q: "Grelina está associada a:", options: ["Estímulo da fome","Saciedade plena","Sono REM","Memória semântica"], answer: 0, explain: "Hormônio do estômago — aumenta apetite." },
        { q: "Teoria dos instintos foi criticada por:", options: ["Explicar tudo com rótulos sem testar","Excesso de rigor experimental","Foco em cognição","Rejeitar biologia"], answer: 0, explain: "McDougal rotulava comportamentos como instintos ad hoc." },
        { q: "Motivação extrínseca depende de:", options: ["Recompensas e punições externas","Prazer puro interno","Instinto irreprimível","Reflexo medular"], answer: 0, explain: "Nota, dinheiro, aprovação social." },
        { q: "Comer emocional relaciona fome a:", options: ["Estresse e emoção","Apenas hipotálamo","Memória procedural","Percepção gestáltica"], answer: 0, explain: "Fatores psicológicos além da necessidade calórica." },
        { q: "Segundo Maslow, necessidades básicas incluem:", options: ["Fisiológicas e segurança","Apenas autorrealização","Somente estima social","Exclusivamente sexual"], answer: 0, explain: "Base da pirâmide: sobrevivência e segurança." }
      ]
    }
  },
  {
    id: 7,
    title: "Emoção",
    icon: "❤️",
    color: "#ef4444",
    description: "Componentes e teorias de James-Lange a Schachter",
    lessons: [
      {
        title: "Componentes da Emoção",
        content: `<p>Três componentes interligados:</p>
        <ul>
          <li><strong>Fisiológico:</strong> ativação do SNS, adrenalina, expressão facial</li>
          <li><strong>Cognitivo:</strong> avaliação e interpretação do evento</li>
          <li><strong>Comportamental:</strong> ação, expressão, comunicação</li>
        </ul>
        <p>Emoções básicas (Ekman): alegria, tristeza, medo, raiva, surpresa, nojo.</p>`
      },
      {
        title: "Teoria de James-Lange",
        content: `<p><strong>James-Lange:</strong> estímulo → resposta corporal → emoção.</p>
        <p>"Choramos porque temos medo" invertido: "Temos medo porque trememos."</p>
        <p>Ênfase no feedback corporal como fonte da experiência emocional.</p>`
      },
      {
        title: "Teoria Cannon-Bard",
        content: `<p><strong>Cannon-Bard:</strong> estímulo → ativação simultânea de emoção e resposta fisiológica.</p>
        <p>Crítica a James-Lange: mudanças corporais são lentas e pouco diferenciadas.</p>
        <p>Tálamo mediaria experiência emocional e resposta.</p>`
      },
      {
        title: "Teoria dos Dois Fatores",
        content: `<p><strong>Schachter-Singer:</strong> emoção = arousal fisiológico + <strong>rotulação cognitiva</strong>.</p>
        <p>Experimento da epinefrina: participantes rotulavam arousal conforme contexto social.</p>
        <div class="highlight">Sem interpretação cognitiva, arousal permanece difuso.</div>`
      },
      {
        title: "Emoção e Cognição",
        content: `<p><strong>Teoria da avaliação (Lazarus):</strong> avaliação primária e secundária precede emoção.</p>
        <p>Emoções influenciam memória, atenção e decisão.</p>
        <p>Inteligência emocional (Goleman): reconhecer e regular emoções.</p>`
      }
    ],
    quiz: [
      { q: "James-Lange propõe que emoção surge de:", options: ["Percepção de resposta corporal","Apenas pensamento puro","Instinto irreprimível","Memória de longo prazo"], answer: 0, explain: "Feedback fisiológico gera experiência emocional." },
      { q: "Cannon-Bard defende que emoção e fisiologia são:", options: ["Simultâneas","Sempre sequenciais com corpo primeiro","Inexistentes","Puramente cognitivas"], answer: 0, explain: "Ativação paralela, não causal corporal → emoção." },
      { q: "Schachter-Singer enfatiza dois fatores:", options: ["Arousal + rotulação cognitiva","Fome + sono","Percepção + sensação","Memória + linguagem"], answer: 0, explain: "Interpretamos arousal conforme contexto." },
      { q: "Ekman estudou principalmente:", options: ["Expressões faciais universais","Motivação de fome","Sono REM","Heurísticas de decisão"], answer: 0, explain: "Emoções básicas reconhecíveis transculturalmente." },
      { q: "Sistema nervoso simpático ativa-se em:", options: ["Emoções de alerta e medo","Sono profundo NREM","Digestão em repouso","Memória procedural"], answer: 0, explain: "Fight-or-flight — prepara corpo para ação." }
    ],
    matching: [
      { term: "James-Lange", match: "Corpo antes da emoção" },
      { term: "Cannon-Bard", match: "Ativação simultânea" },
      { term: "Schachter-Singer", match: "Arousal + cognição" },
      { term: "Ekman", match: "Expressões universais" },
      { term: "Lazarus", match: "Avaliação cognitiva" }
    ],
    boss: {
      name: "Tempestade Emocional",
      desc: "Prove domínio das teorias da emoção!",
      questions: [
        { q: "Experimento da epinefrina (Schachter) mostrou:", options: ["Contexto influencia rotulação emocional","Emoção não existe","Corpo não reage","Memória é perfeita"], answer: 0, explain: "Mesmo arousal, rótulos diferentes conforme situação." },
        { q: "Crítica a James-Lange: respostas corporais são:", options: ["Pouco específicas por emoção","Sempre únicas por emoção","Inexistentes","Puramente cognitivas"], answer: 0, explain: "Cannon: coração acelera em medo e raiva — pouco discriminativo." },
        { q: "Avaliação primária (Lazarus) pergunta:", options: ["Isso é relevante para mim?","Qual cor do objeto?","Quantos itens na lista?","Qual hora do sono?"], answer: 0, explain: "Relevância pessoal do evento." },
        { q: "Adrenalina relaciona-se a:", options: ["Ativação fisiológica emocional","Sono profundo","Memória semântica","Percepção de profundidade"], answer: 0, explain: "Hormônio do estresse — aumenta arousal." },
        { q: "Componente comportamental da emoção inclui:", options: ["Expressão facial e ação","Apenas pensamento","Somente sonho REM","Exclusivamente memória"], answer: 0, explain: "Como emoção se manifesta externamente." }
      ]
    }
  },
  {
    id: 8,
    title: "Pensamento",
    icon: "💡",
    color: "#10b981",
    description: "Heurísticas, vieses, decisão e linguagem",
    lessons: [
      {
        title: "Conceito de Pensamento",
        content: `<p><strong>Pensamento</strong> = manipulação mental de informações: conceitos, imagens, problemas.</p>
        <p>Formas: raciocínio dedutivo, indutivo, criativo, crítico.</p>
        <p>Sternberg: psicologia cognitiva estuda processos de pensamento.</p>`
      },
      {
        title: "Heurísticas",
        content: `<p><strong>Heurísticas</strong> = atalhos mentais para decisões rápidas (nem sempre ótimas).</p>
        <ul>
          <li><strong>Disponibilidade:</strong> julgar pela facilidade de lembrar exemplos</li>
          <li><strong>Representatividade:</strong> julgar pela semelhança a protótipo</li>
          <li><strong>Ancoragem:</strong> fixar em primeira informação</li>
        </ul>`
      },
      {
        title: "Vieses Cognitivos",
        content: `<p><strong>Viés de confirmação:</strong> buscar informação que confirma crenças.</p>
        <p><strong>Excesso de confiança:</strong> superestimar acerto de julgamentos.</p>
        <p><strong>Aversão à perda (Kahneman & Tversky):</strong> perdas pesam mais que ganhos equivalentes.</p>
        <div class="highlight">Vieses são sistemáticos — prospect theory revolucionou economia comportamental.</div>`
      },
      {
        title: "Resolução de Problemas",
        content: `<p>Etapas: identificar → representar → estratégias (algoritmo, heurística, insight).</p>
        <p><strong>Fixação funcional:</strong> ver objeto só no uso habitual.</p>
        <p><strong>Problema das 9 pontos:</strong> pensar "fora da caixa".</p>
        <p>Tomada de decisão: racional limitada (Simon) vs modelo econômico clássico.</p>`
      },
      {
        title: "Linguagem e Pensamento",
        content: `<p><strong>Linguagem</strong> medeia pensamento — debate Sapir-Whorf (relatividade linguística).</p>
        <p>Componentes: fonologia, morfologia, sintaxe, semântica, pragmática.</p>
        <p>Aquisição: Chomsky (gramática universal) vs aprendizagem comportamental.</p>
        <p>Pensamento e linguagem são profundamente entrelaçados na cognição humana.</p>`
      }
    ],
    quiz: [
      { q: "Heurística da disponibilidade usa:", options: ["Facilidade de lembrar exemplos","Apenas lógica formal","Instinto animal","Memória procedural"], answer: 0, explain: "Eventos fáceis de recordar parecem mais frequentes." },
      { q: "Kahneman e Tversky propuseram:", options: ["Teoria da perspectiva","Behaviorismo radical","Psicanálise freudiana","Gestalt de percepção"], answer: 0, explain: "Prospect theory — decisão sob risco e vieses." },
      { q: "Viés de confirmação leva a:", options: ["Buscar só evidências favoráveis","Decisão sempre racional","Memória perfeita","Percepção neutra"], answer: 0, explain: "Ignoramos informação contrária às crenças." },
      { q: "Fixação funcional impede:", options: ["Usar objeto de forma nova","Ver cores corretamente","Dormir adequadamente","Sentir emoções"], answer: 0, explain: "Martelo só como martelo — não como peso." },
      { q: "Chomsky defendeu:", options: ["Gramática universal inata","Aprendizagem só por reforço","Inexistência de linguagem","Pensamento sem linguagem"], answer: 0, explain: "Competência linguística inata — crítica ao behaviorismo." }
    ],
    matching: [
      { term: "Heurística", match: "Atalho mental" },
      { term: "Ancoragem", match: "Fixar em referência" },
      { term: "Kahneman", match: "Teoria da perspectiva" },
      { term: "Chomsky", match: "Gramática universal" },
      { term: "Insight", match: "Solução súbita" }
    ],
    boss: {
      name: "Mente Suprema",
      desc: "O chefe final! Domine todo o semestre!",
      questions: [
        { q: "Aversão à perda significa que:", options: ["Perdas doem mais que ganhos equivalentes","Ganhos sempre superam perdas","Não há decisão racional","Emoção não importa"], answer: 0, explain: "Assimetria na avaliação de resultados." },
        { q: "Racionalidade limitada (Simon) reconhece:", options: ["Capacidade cognitiva finita","Perfeição nas decisões","Ausência de vieses","Pensamento infinito"], answer: 0, explain: "Satisficing — escolher opção boa o suficiente." },
        { q: "Hipótese Sapir-Whorf sugere que:", options: ["Linguagem influencia pensamento","Pensamento é independente de linguagem","Emoção não existe","Memória é infinita"], answer: 0, explain: "Relatividade linguística — debate contínuo." },
        { q: "Sternberg estuda principalmente:", options: ["Psicologia cognitiva","Neuroanatomia pura","Psicanálise kleiniana","Behaviorismo skinneriano"], answer: 0, explain: "Psicologia cognitiva — pensamento e inteligência." },
        { q: "Tomada de decisão envolve:", options: ["Julgamento, risco e vieses","Apenas reflexos","Somente instinto","Exclusivamente sono"], answer: 0, explain: "Processo cognitivo complexo com erros sistemáticos." }
      ]
    }
  }
];

const ACHIEVEMENTS = [
  { id: "first_lesson", icon: "📖", title: "Primeira Lição", desc: "Complete sua primeira lição", xp: 0 },
  { id: "first_quiz", icon: "✅", title: "Quizzer", desc: "Complete seu primeiro quiz", xp: 0 },
  { id: "first_boss", icon: "⚔️", title: "Caçador de Chefes", desc: "Derrote seu primeiro chefe", xp: 0 },
  { id: "unit1", icon: "🔬", title: "Cientista", desc: "Complete a Unidade 1", xp: 0 },
  { id: "unit4", icon: "💾", title: "Mnemonista", desc: "Complete a Unidade 4", xp: 0 },
  { id: "unit8", icon: "🏆", title: "Mestre Cognitivo", desc: "Complete todas as 8 unidades", xp: 0 },
  { id: "streak5", icon: "🔥", title: "Em Chamas", desc: "Acerte 5 questões seguidas", xp: 0 },
  { id: "streak10", icon: "💎", title: "Imparável", desc: "Acerte 10 questões seguidas", xp: 0 },
  { id: "no_life_lost", icon: "🛡️", title: "Invencível", desc: "Complete um quiz sem perder vidas", xp: 0 },
  { id: "level5", icon: "⭐", title: "Veterano", desc: "Alcance o nível 5", xp: 0 },
  { id: "level10", icon: "👑", title: "Lenda Mental", desc: "Alcance o nível 10", xp: 0 },
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

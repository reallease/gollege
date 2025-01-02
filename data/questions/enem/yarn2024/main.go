package yarn2024

import (
	questions "github.com/reallease/gollege/data/questions/enem"
)

var Languages = []questions.Question{
	{
		// ID: "ENEM 2019 - 01",
		// ID: "enem2019_q01",
		ID:   "enem2023_q01", // acertei
		Name: "ENEM",
		Year: "2023",
		Files: []string{
			"https://d23vy2bv3rsfba.cloudfront.net/questoes_imagens/0_acc99be78fb25853ea191188dd376876_10317920.jpg.png",
		},
		Statement: "Ao retratar o ambiente de trabalho em um escritório, esse cartum tem por objetivo",
		Subject:   "english",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) criticar um padrão de vestimenta.",
				Correct:   false,
			},

			{
				Statement: "b) destacar a falta de diversidade.",
				Correct:   true, // is correct
			},

			{
				Statement: "c) indicar um modo de interação.",
				Correct:   false,
			},

			{
				Statement: "d) elogiar um modelo de organização.",
				Correct:   false,
			},

			{
				Statement: "e) salientar o espírito de cooperação.",
				Correct:   false,
			},
		},

		Credits:   "Disponível em: www.cartoonstock.com. Acesso em: 25 out. 2021.",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q02", // acertei
		Name:      "ENEM",
		Year:      "2023",
		Question:  "No man is an island,\n Entire of itself;\n Every man is a piece of the continent,\n A part of the main.\n [...] \n Any man's death diminishes me,\nBecause I am involved in mankind.",
		Statement: "Nesse poema, a expressão “No man is an island” ressalta o(a)",
		Subject:   "english",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) medo da morte.",
				Correct:   false,
			},

			{
				Statement: "b) ideia de conexão.",
				Correct:   true,
			},

			{
				Statement: "c) 	conceito de solidão.",
				Correct:   false,
			},

			{
				Statement: "d) risco de devastação.",
				Correct:   false,
			},

			{
				Statement: "e) necessidade de empatia.",
				Correct:   false,
			},
		},
		Credits:   "DONNE, J. The Works of John Donne. Londres: John W. Parker, 1839 (fragmento).",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:   "enem2023_q03", // acertei
		Name: "ENEM",
		Year: "2023",
		Files: []string{
			"https://d23vy2bv3rsfba.cloudfront.net/questoes_imagens/0_0086898809eb90513ad6eda8143a5fee_10317970.jpg.png",
		},
		Statement: "Esse cartaz de campanha sugere que",
		Subject:   "english",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) os lixões precisam de ampliação.",
				Correct:   false,
			},

			{
				Statement: "b) o desperdício degrada o ambiente.",
				Correct:   true,
			},

			{
				Statement: "c) os mercados doam alimentos perecíveis.",
				Correct:   false,
			},

			{
				Statement: "d) a desnutrição compromete o raciocínio.",
				Correct:   false,
			},

			{
				Statement: "e) as residências carecem de refrigeradores.",
				Correct:   false,
			},
		},

		Credits:   "Disponível em: https://mir-s3-cdn-cf.behance.net. Acesso em: 29 out. 2021 (adaptado).",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q04", //
		Name:      "ENEM",
		Year:      "2023",
		Question:  "Things We Carry on the Sea\n\nWe carry tears in our eyes: good-bye father, good-bye mother  \nWe carry soi in smal bags: may home never fade in our hearts  \nWe carry carnage of mining. droughts, floods, genocides  \nWe carry dust of our families and neighbors incinerated in mushroom clouds  \nWe carry our islands sinking under the sea  \nWe carry our hands, feet, bones, hearts and best minds for a new life  \nWe carry diplomas: medicine, engineer, nurse, education, math, poetry, even if they mean nothing to the other shore  \nWe carry railroads, plantations, laundromats. bodegas, taco trucks, farms, factories, nursing homes, hospitals, schools, temples… built on our ancestors’ backs  \nWe carry old homes along the spine, new dreams in our chests  \nWe carry yesterday, today and tomorrow  \nWe’re orphans of the wars forced upon us  \nWe’re refugees of the sea rising from industrial wastes  \nAnd we carry our mother tongues  \n\\[…\\]  \nAs we drift… in our rubber boats… from shore… to shore…. to shore…",
		Statement: "Ao retratar a trajetória de refugiados, o poema recorre à 	imagem de viagem marítima para destacar o(a)",
		Subject:   "english",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) risco de choques culturais.",
				Correct:   false,
			},

			{
				Statement: "b) impacto do ensino de história.",
				Correct:   false,
			},

			{
				Statement: "c) importância da luta ambiental.",
				Correct:   false,
			},

			{
				Statement: "d) existência de experiências plurais.",
				Correct:   true,
			},

			{
				Statement: "e) necessidade de capacitação profissional.",
				Correct:   false,
			},
		},
		Credits:   "PING, W. Disponível em: https://poets.org. Acesso em: 1 jun. 2023 (fragmento).",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q05",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "Spanglish\n\n pues estoy creando Spanglish\n bi-cultural systems\n scientific lexicographical\n inter-textual integrations\n two expressions\n existentially wired\n two dominant languages\n continentally abrazándose\n in colloquial combate\n imperio spanglish emerges\n sobre territorio bi-lingual\n las novelas mexicanas\n mixing with radiorocknroll\n immigrant/migrant\n nasal mispronouncements\n hip-hop, street salsa, spanish pop\n standard english classroom\n with computer technicalities\n spanglish is literally perfect",
		Statement: "Nesse poema de Tato Laviera, o eu lírico destaca uma",
		Subject:   "english",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) convergência linguístico-cultural.",
				Correct:   true,
			},

			{
				Statement: "b) característica histórico-cultural.",
				Correct:   false,
			},

			{
				Statement: "c) tendência estilístico-literária.",
				Correct:   false,
			},

			{
				Statement: "d) discriminação cultural.",
				Correct:   false,
			},

			{
				Statement: "e) censura musical.",
				Correct:   false,
			},
		},

		Credits:   "LAVIERA, T. Benedición: The Complete Poetry of Tato Laviera. Houston: Arte Público Press, 2014 (fragmento).",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q06", // err
		Name:      "ENEM",
		Year:      "2023",
		Question:  "Me niego rotundamente\nA negar mi voz,\nMi sangre y mi piel.\nY me niego rotundamente\nA dejar de ser yo,\nA dejar de sentirme bien\nCuando miro mi rostro en el es\nCon mi boca\nRotundamente grande\nY mi nariz\nRotundamente hermosa,\nY mis dientes\nRotundamente blancos,\nY mi piel valientemente negra.\nY me niego categóricamente\nA dejar de hablar\nMi lengua, mi acento y mi historia\nY me niego absolutamente\nA ser parte de los que callan,\nDe los que temen,\nDe los que lloran\nPorque me acepto\nRotundamente libre,\nRotundamente negra,\nRotundamente hermosa.\nCAMPELL BARR, S. Disponível em:\nhttps.negracubanateniaqueser.com\nAcesso em 25 out. 2021.\nPara enfatizar caracteristicas e atitudes que reforçam a identidade da mulher negra, o poema da escritora costarriquenha apresenta:",
		Statement: "Para enfatizar características e atitudes que reforçam a identidade da mulher negra, o poema da escritora costarriquenha apresenta",
		Subject:   "spanish",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) advérbios como “rotundamente” e “categóricamente”.",
				Correct:   true,
			},

			{
				Statement: "b) verbos reflexivos como “me niego” e “me acepto”.",
				Correct:   false,
			},

			{
				Statement: "c) adjetivos como “grande” e “hermosa”.",
				Correct:   false,
			},

			{
				Statement: "d) substantivos como “sangre” e “piel”.",
				Correct:   false,
			},

			{
				Statement: "e) adjetivos possessivos como “mi” e “mis”.",
				Correct:   false,
			},
		},

		Credits:   "CAMPBELL BARR, S. Disponível em: https://negracubanateniaqueser.com. Acesso em: 25 out. 2021",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q07", // talvez formatação dos paragrafos!
		Name:      "ENEM",
		Year:      "2023",
		Question:  "Técnicas de manipulación y el resultado\n\nManipular es sembrar en la conciencia y en la mente de la gente ideas, actitudes, conceptos y aspiraciones — incluso falsas e inmorales — que sirvan a los objetivos de sus manipuladores.\nManipular es una de las primeras cosas que aprendemos en la vida. A muy temprana edad, los bebés descubren el poder del Ilanto, el berrinche, los pataleos, la risa o alguna “gracia” como recursos para demandar atención, exigir comida, pedir ayuda o simplemente mantener ocupada a la gente. Nuestras actitudes de adultos reflejan lo mucho o poco que algunos maduraron, procesaron y rebasaron ese periodo.\nPara que exista un manipulador, debe haber una base de ciudadanos indefensos, dóciles, desinformados. El manipulador es celoso, a veces casi paranoico; no admite cuestionamientos ni quiere que nadie ocupe su espacio, sabe que su vigencia depende de presencia controladora. Todos los días, hay que marcar la línea de discurso, incidir en el debate. El ridículo vale la pena si con ello se logra una cortina de humo.",
		Statement: "Nesse texto, a expressão “cortina de humo” revela que o manipulador",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "a) amadurece tardiamente.",
				Correct:   false,
			},

			{
				Statement: "b) busca mascarar a verdade.",
				Correct:   false,
			},

			{
				Statement: "c) rejeita questionamentos alheios.",
				Correct:   false,
			},

			{
				Statement: "d) aproxima-se de pessoas indefesas.",
				Correct:   false,
			},

			{
				Statement: "e) faz-se presente de forma controladora.",
				Correct:   false,
			},
		},

		Credits:   "Disponível em: www.forbes.com.mx. Acesso em: 7 out. 2021 (adaptado).",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q07",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q08",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q09",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q10",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q11",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q12",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q13",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q14",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q15",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q16",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q17",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q18",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q19",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q20",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q21",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q22",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q23",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q24",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q25",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q26",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q27",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q28",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q29",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
	{
		ID:        "enem2023_q30",
		Name:      "ENEM",
		Year:      "2023",
		Question:  "",
		Statement: "",
		Subject:   "",

		Alternatives: []questions.Alternative{
			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},

			{
				Statement: "",
				Correct:   false,
			},
		},

		Credits:   "",
		Copyright: "https://download.inep.gov.br/enem/provas_e_gabaritos/2023_PV_impresso_D1_CD2.pdf",
	},
}

var Natural_sciences = []questions.Question{
	{},
}

var Humans_sciences = []string{
	1: "1",
}

var Math = []questions.Question{
	{},
}

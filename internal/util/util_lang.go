package util

var (
	currentLanguage = "en-US"

	errorMessages = map[string]map[string]map[string]map[string]string{
		"en-US": {
			"LoginUseCase": {
				"UserNotFound": {
					"Title":  "User Not Found",
					"Detail": "The user with the provided credentials could not be found.",
				},
				"UserNotActive": {
					"Title":  "User Not Active",
					"Detail": "The user account is not active. Please contact support.",
				},
				"InvalidCredentials": {
					"Title":  "Invalid Credentials",
					"Detail": "The email or password provided is incorrect.",
				},
				"JWTError": {
					"Title":  "JWT Error",
					"Detail": "An error occurred while generating the JWT token.",
				},
			},
		},
		"pt-BR": {
			"LoginUseCase": {
				"UserNotFound": {
					"Title":  "Usuário Não Encontrado",
					"Detail": "Não foi possível encontrar o usuário com as credenciais fornecidas.",
				},
				"UserNotActive": {
					"Title":  "Usuário Não Ativo",
					"Detail": "A conta do usuário não está ativa. Entre em contato com o suporte.",
				},
				"InvalidCredentials": {
					"Title":  "Credenciais Inválidas",
					"Detail": "O e-mail ou a senha fornecida está incorreta.",
				},
				"JWTError": {
					"Title":  "Erro JWT",
					"Detail": "Ocorreu um erro ao gerar o token JWT.",
				},
			},
		},
		"fr-FR": {
			"LoginUseCase": {
				"UserNotFound": {
					"Title":  "Utilisateur Non Trouvé",
					"Detail": "L'utilisateur avec les identifiants fournis n'a pas été trouvé.",
				},
				"UserNotActive": {
					"Title":  "Utilisateur Inactif",
					"Detail": "Le compte utilisateur n'est pas actif. Veuillez contacter le support.",
				},
				"InvalidCredentials": {
					"Title":  "Identifiants Invalide",
					"Detail": "L'email ou le mot de passe fourni est incorrect.",
				},
				"JWTError": {
					"Title":  "Erreur JWT",
					"Detail": "Une erreur est survenue lors de la génération du jeton JWT.",
				},
			},
		},
		"es-ES": {
			"LoginUseCase": {
				"UserNotFound": {
					"Title":  "Usuario No Encontrado",
					"Detail": "No se pudo encontrar al usuario con las credenciales proporcionadas.",
				},
				"UserNotActive": {
					"Title":  "Usuario No Activo",
					"Detail": "La cuenta del usuario no está activa. Por favor contacte con el soporte.",
				},
				"InvalidCredentials": {
					"Title":  "Credenciales Inválidas",
					"Detail": "El correo electrónico o la contraseña proporcionados son incorrectos.",
				},
				"JWTError": {
					"Title":  "Error JWT",
					"Detail": "Ocurrió un error al generar el token JWT.",
				},
			},
		},
		"zh-CN": {
			"LoginUseCase": {
				"UserNotFound": {
					"Title":  "用户未找到",
					"Detail": "无法找到提供的凭据对应的用户。",
				},
				"UserNotActive": {
					"Title":  "用户未激活",
					"Detail": "该用户帐户未激活。请联系支持团队。",
				},
				"InvalidCredentials": {
					"Title":  "无效的凭据",
					"Detail": "提供的电子邮件或密码错误。",
				},
				"JWTError": {
					"Title":  "JWT 错误",
					"Detail": "生成JWT令牌时发生错误。",
				},
			},
		},
	}
)

func SetLanguage(lang string) {
	currentLanguage = lang
}

func GetLanguage() string {
	return currentLanguage
}

func GetErrorMessage(module, key, messageType string) string {
	if msg, ok := errorMessages[currentLanguage][module][key]; ok {
		if detail, ok := msg[messageType]; ok {
			return detail
		}
	}
	return errorMessages["en-US"][module][key]["Detail"]
}

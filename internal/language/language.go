package language

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"

var (
	currentLanguage = "en-US"

	errorMessages = map[string]map[string]map[string]map[string]string{
		"en-US": {
			"CommonErrors": {
				"JsonBindingError": {
					"Title":  "Failed to process request",
					"Detail": "Could not parse the request body. Please check the JSON format sent.",
				},
				"MissingUserID": {
					"Title":  "Missing User ID",
					"Detail": "User ID is required to access this resource.",
				},
				"InvalidUserID": {
					"Title":  "Invalid User ID",
					"Detail": "A valid user ID must be provided.",
				},
			},
			"CreateMovieUseCase": {
				"MovieAlreadyExists": {
					"Title":  "Movie already exists",
					"Detail": "A movie with the same external ID already exists. Please check the external ID and try again.",
				},
				"ErrorFetchingExistingMovie": {
					"Title":  "Error fetching existing movie",
					"Detail": "An error occurred while checking if the movie already exists.",
				},
				"ErrorSavingPoster": {
					"Title":  "Error saving poster",
					"Detail": "An error occurred while saving the movie poster. Please try again later.",
				},
				"ErrorCreatingMovie": {
					"Title":  "Error creating movie",
					"Detail": "An error occurred while creating the movie in the database.",
				},
			},
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
				"HMAC": {
					"Title":  "HMAC Error",
					"Detail": "An error occurred while hashing the email with HMAC.",
				},
			},
			"CreateBrandUseCase": {
				"BrandAlreadyExists": {
					"Title":  "Conflict",
					"Detail": "A brand with the same name already exists in the system.",
				},
				"ErrorFetchingBrand": {
					"Title":  "Error fetching existing brand",
					"Detail": "An unexpected error occurred while verifying existing brand data.",
				},
				"InvalidBrandData": {
					"Title":  "Invalid Brand Data",
					"Detail": "The provided brand data is invalid. Please review the input fields.",
				},
				"ErrorSavingLogo": {
					"Title":  "Error saving logo",
					"Detail": "We encountered an issue while saving the brand's logo. Please try again later.",
				},
				"ErrorCreatingBrand": {
					"Title":  "Error creating brand",
					"Detail": "Something went wrong while creating the brand. Please contact support if the issue persists.",
				},
			},
			"CreateUserUseCase": {
				"EmailHMACError": {
					"Title":  "Error hashing email",
					"Detail": "An error occurred while hashing the email with HMAC.",
				},
				"EmailAlreadyExists": {
					"Title":  "Email already exists",
					"Detail": "The provided email is already registered.",
				},
				"EmailCheckError": {
					"Title":  "Error checking user email existence",
					"Detail": "An error occurred while checking if the email already exists.",
				},
				"UsernameAlreadyExists": {
					"Title":  "Username already exists",
					"Detail": "The provided username is already taken.",
				},
				"UsernameCheckError": {
					"Title":  "Error checking user name existence",
					"Detail": "An error occurred while checking if the username already exists.",
				},
				"NewLoginError": {
					"Title":  "Error creating login",
					"Detail": "An error occurred while creating the login entity.",
				},
				"HashEmailError": {
					"Title":  "Error encrypting email",
					"Detail": "An error occurred while encrypting the email address.",
				},
				"HashPasswordError": {
					"Title":  "Error encrypting password",
					"Detail": "An error occurred while encrypting the password.",
				},
				"NewUserError": {
					"Title":  "Error creating user",
					"Detail": "An error occurred while creating the new user.",
				},
				"SaveUserError": {
					"Title":  "Error creating new user",
					"Detail": "An error occurred while saving the new user to the database.",
				},
			},
			"AddBrandsListUseCase": {
				"UserNotFound": {
					"Title":  "User Not Found",
					"Detail": "The user with the provided ID could not be found.",
				},
				"UserNotActive": {
					"Title":  "User Not Active",
					"Detail": "The user account is not active. Please contact support.",
				},
				"UserNotAdmin": {
					"Title":  "User Not Admin",
					"Detail": "The user does not have admin privileges.",
				},
				"ListNotFound": {
					"Title":  "List Not Found",
					"Detail": "The list with the provided ID could not be found.",
				},
				"InvalidListType": {
					"Title":  "Invalid List Type",
					"Detail": "The list type must be 'brand'.",
				},
				"BrandAlreadyInList": {
					"Title":  "Brand Already In List",
					"Detail": "The brand with the provided ID already exists in the list.",
				},
				"ErrorFetchingBrands": {
					"Title":  "Error Fetching Brands",
					"Detail": "An error occurred while fetching the brands.",
				},
				"ErrorAddingBrands": {
					"Title":  "Error Adding Brands",
					"Detail": "An error occurred while adding the brands to the list.",
				},
			},
			"AddMoviesListUseCase": {
				"UserNotFound": {
					"Title":  "User Not Found",
					"Detail": "The user with the provided ID could not be found.",
				},
				"UserNotActive": {
					"Title":  "User Not Active",
					"Detail": "The user's account is not active. Please contact support.",
				},
				"UserNotAdmin": {
					"Title":  "Access Denied",
					"Detail": "The user does not have admin privileges.",
				},
				"ListNotFound": {
					"Title":  "List Not Found",
					"Detail": "The list with the provided ID could not be retrieved.",
				},
				"InvalidListType": {
					"Title":  "Invalid List Type",
					"Detail": "The list type must be 'movie'.",
				},
				"MovieAlreadyInList": {
					"Title":  "Movie Already in List",
					"Detail": "The movie is already present in the list.",
				},
				"ErrorFetchingMovies": {
					"Title":  "Error Fetching Movies",
					"Detail": "An error occurred while fetching movies with the provided IDs.",
				},
				"ErrorAddingMovies": {
					"Title":  "Error Adding Movies",
					"Detail": "An error occurred while adding movies to the list.",
				},
			},
			"AuthMiddleware": {
				"UnauthorizedHeader": {
					"Title":  "Missing Authorization Header",
					"Detail": "Authorization header is required",
				},
				"UnauthorizedBearer": {
					"Title":  "Invalid Authorization Format",
					"Detail": "Authorization header must be in the format 'Bearer <token>'",
				},
				"UnauthorizedTokenParse": {
					"Title":  "Unexpected signing method",
					"Detail": "Unexpected signing method",
				},
				"UnauthorizedInvalidToken": {
					"Title":  "Invalid Token",
					"Detail": "Token could not be parsed or is invalid",
				},
				"UnauthorizedToken": {
					"Title":  "Invalid Token",
					"Detail": "Token is not valid",
				},
			},
		},
		"pt-BR": {
			"CommonErrors": {
				"JsonBindingError": {
					"Title":  "Erro ao processar requisição",
					"Detail": "Não foi possível interpretar o corpo da requisição. Verifique o formato do JSON enviado.",
				},
				"MissingUserID": {
					"Title":  "ID do Usuário ausente",
					"Detail": "O ID do usuário é obrigatório para acessar este recurso.",
				},
				"InvalidUserID": {
					"Title":  "ID do Usuário inválido",
					"Detail": "É necessário fornecer um ID de usuário válido.",
				},
			},
			"CreateMovieUseCase": {
				"MovieAlreadyExists": {
					"Title":  "Filme já existe",
					"Detail": "Já existe um filme com o mesmo ID externo. Verifique o ID externo e tente novamente.",
				},
				"ErrorFetchingExistingMovie": {
					"Title":  "Erro ao buscar filme existente",
					"Detail": "Ocorreu um erro ao verificar se o filme já existe.",
				},
				"ErrorSavingPoster": {
					"Title":  "Erro ao salvar o pôster",
					"Detail": "Ocorreu um erro ao salvar o pôster do filme. Tente novamente mais tarde.",
				},
				"ErrorCreatingMovie": {
					"Title":  "Erro ao criar filme",
					"Detail": "Ocorreu um erro ao salvar o filme no banco de dados.",
				},
			},
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
				"HMAC": {
					"Title":  "Erro HMAC",
					"Detail": "Ocorreu um erro ao fazer o hash do e-mail com HMAC.",
				},
			},
			"CreateBrandUseCase": {
				"BrandAlreadyExists": {
					"Title":  "Conflito",
					"Detail": "Já existe uma marca com o mesmo nome no sistema.",
				},
				"ErrorFetchingBrand": {
					"Title":  "Erro ao verificar marca existente",
					"Detail": "Ocorreu um erro inesperado ao verificar os dados da marca.",
				},
				"InvalidBrandData": {
					"Title":  "Dados inválidos da marca",
					"Detail": "Os dados fornecidos para a marca são inválidos. Verifique os campos e tente novamente.",
				},
				"ErrorSavingLogo": {
					"Title":  "Erro ao salvar logotipo",
					"Detail": "Ocorreu um problema ao salvar o logotipo da marca. Tente novamente mais tarde.",
				},
				"ErrorCreatingBrand": {
					"Title":  "Erro ao criar marca",
					"Detail": "Algo deu errado ao tentar criar a marca. Caso o problema persista, entre em contato com o suporte.",
				},
			},
			"CreateUserUseCase": {
				"EmailHMACError": {
					"Title":  "Erro ao gerar hash do e-mail",
					"Detail": "Ocorreu um erro ao aplicar o HMAC no endereço de e-mail.",
				},
				"EmailAlreadyExists": {
					"Title":  "E-mail já cadastrado",
					"Detail": "O e-mail informado já está registrado em nossa base de dados.",
				},
				"EmailCheckError": {
					"Title":  "Erro ao verificar existência do e-mail",
					"Detail": "Ocorreu um erro ao verificar se o e-mail já está cadastrado.",
				},
				"UsernameAlreadyExists": {
					"Title":  "Nome de usuário já em uso",
					"Detail": "O nome de usuário informado já está sendo utilizado por outro usuário.",
				},
				"UsernameCheckError": {
					"Title":  "Erro ao verificar existência do nome de usuário",
					"Detail": "Ocorreu um erro ao verificar se o nome de usuário já está cadastrado.",
				},
				"NewLoginError": {
					"Title":  "Erro ao criar login",
					"Detail": "Ocorreu um erro ao criar a entidade de login.",
				},
				"HashEmailError": {
					"Title":  "Erro ao criptografar e-mail",
					"Detail": "Ocorreu um erro ao criptografar o endereço de e-mail.",
				},
				"HashPasswordError": {
					"Title":  "Erro ao criptografar senha",
					"Detail": "Ocorreu um erro ao criptografar a senha do usuário.",
				},
				"NewUserError": {
					"Title":  "Erro ao criar usuário",
					"Detail": "Ocorreu um erro ao instanciar o novo usuário.",
				},
				"SaveUserError": {
					"Title":  "Erro ao salvar usuário",
					"Detail": "Ocorreu um erro ao salvar o novo usuário no banco de dados.",
				},
			},
			"AddBrandsListUseCase": {
				"UserNotFound": {
					"Title":  "Usuário Não Encontrado",
					"Detail": "Não foi possível encontrar o usuário com o ID fornecido.",
				},
				"UserNotActive": {
					"Title":  "Usuário Não Ativo",
					"Detail": "A conta do usuário não está ativa. Entre em contato com o suporte.",
				},
				"UserNotAdmin": {
					"Title":  "Usuário Não é Administrador",
					"Detail": "O usuário não possui privilégios de administrador.",
				},
				"ListNotFound": {
					"Title":  "Lista Não Encontrada",
					"Detail": "Não foi possível encontrar a lista com o ID fornecido.",
				},
				"InvalidListType": {
					"Title":  "Tipo de Lista Inválido",
					"Detail": "O tipo da lista deve ser 'brand'.",
				},
				"BrandAlreadyInList": {
					"Title":  "Marca Já Na Lista",
					"Detail": "A marca com o ID fornecido já existe na lista.",
				},
				"ErrorFetchingBrands": {
					"Title":  "Erro ao Buscar Marcas",
					"Detail": "Ocorreu um erro ao buscar as marcas.",
				},
				"ErrorAddingBrands": {
					"Title":  "Erro ao Adicionar Marcas",
					"Detail": "Ocorreu um erro ao adicionar as marcas à lista.",
				},
				"ErrorFetchingCombinations": {
					"Title":  "Erro ao Buscar Combinações",
					"Detail": "Ocorreu um erro ao buscar as combinações para a lista.",
				},
			},
			"AddMoviesListUseCase": {
				"UserNotFound": {
					"Title":  "Usuário não encontrado",
					"Detail": "Não foi possível encontrar o usuário com o ID fornecido.",
				},
				"UserNotActive": {
					"Title":  "Usuário não ativo",
					"Detail": "O usuário não está ativo. Por favor, entre em contato com o suporte.",
				},
				"UserNotAdmin": {
					"Title":  "Usuário não é administrador",
					"Detail": "O usuário não tem permissões de administrador para adicionar filmes à lista.",
				},
				"ErrorFetchingList": {
					"Title":  "Erro ao buscar lista",
					"Detail": "Ocorreu um erro ao tentar recuperar a lista com o ID fornecido.",
				},
				"InvalidListType": {
					"Title":  "Tipo de lista inválido",
					"Detail": "O tipo da lista deve ser 'movie'.",
				},
				"MovieAlreadyInList": {
					"Title":  "Filme já na lista",
					"Detail": "O filme com o ID fornecido já está presente na lista.",
				},
				"ErrorFetchingMovies": {
					"Title":  "Erro ao buscar filmes",
					"Detail": "Ocorreu um erro ao tentar buscar os filmes com os IDs fornecidos.",
				},
				"ErrorAddingMovies": {
					"Title":  "Erro ao adicionar filmes",
					"Detail": "Ocorreu um erro ao tentar adicionar os filmes à lista.",
				},
			},
			"AuthMiddleware": {
				"UnauthorizedHeader": {
					"Title":  "Cabeçalho de Autorização Ausente",
					"Detail": "O cabeçalho de autorização é obrigatório",
				},
				"UnauthorizedBearer": {
					"Title":  "Formato de Autorização Inválido",
					"Detail": "O cabeçalho de autorização deve estar no formato 'Bearer <token>'",
				},
				"UnauthorizedTokenParse": {
					"Title":  "Método de Assinatura Inesperado",
					"Detail": "Método de assinatura inesperado",
				},
				"UnauthorizedInvalidToken": {
					"Title":  "Token Inválido",
					"Detail": "O token não pôde ser analisado ou é inválido",
				},
				"UnauthorizedToken": {
					"Title":  "Token Inválido",
					"Detail": "O token não é válido",
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
			"AddBrandsListUseCase": {
				"UserNotFound": {
					"Title":  "Utilisateur Non Trouvé",
					"Detail": "L'utilisateur avec l'ID fourni n'a pas été trouvé.",
				},
				"UserNotActive": {
					"Title":  "Utilisateur Inactif",
					"Detail": "Le compte de l'utilisateur n'est pas actif. Veuillez contacter le support.",
				},
				"UserNotAdmin": {
					"Title":  "Utilisateur Non Administrateur",
					"Detail": "L'utilisateur n'a pas de privilèges d'administrateur.",
				},
				"ListNotFound": {
					"Title":  "Liste Non Trouvée",
					"Detail": "La liste avec l'ID fourni n'a pas été trouvée.",
				},
				"InvalidListType": {
					"Title":  "Type de Liste Invalide",
					"Detail": "Le type de liste doit être 'brand'.",
				},
				"BrandAlreadyInList": {
					"Title":  "Marque Déjà Dans La Liste",
					"Detail": "La marque avec l'ID fourni existe déjà dans la liste.",
				},
				"ErrorFetchingBrands": {
					"Title":  "Erreur de Récupération des Marques",
					"Detail": "Une erreur s'est produite lors de la récupération des marques.",
				},
				"ErrorAddingBrands": {
					"Title":  "Erreur d'Ajout de Marques",
					"Detail": "Une erreur s'est produite lors de l'ajout des marques à la liste.",
				},
				"ErrorFetchingCombinations": {
					"Title":  "Erreur de Récupération des Combinaisons",
					"Detail": "Une erreur s'est produite lors de la récupération des combinaisons pour la liste.",
				},
			},
			"AddMoviesListUseCase": {
				"UserNotFound": {
					"Title":  "Utilisateur non trouvé",
					"Detail": "L'utilisateur avec l'ID fourni n'a pas pu être trouvé.",
				},
				"UserNotActive": {
					"Title":  "Utilisateur inactif",
					"Detail": "Le compte de l'utilisateur n'est pas actif. Veuillez contacter le support.",
				},
				"UserNotAdmin": {
					"Title":  "Accès refusé",
					"Detail": "L'utilisateur ne possède pas les privilèges d'administrateur.",
				},
				"ListNotFound": {
					"Title":  "Liste non trouvée",
					"Detail": "La liste avec l'ID fourni n'a pas pu être récupérée.",
				},
				"InvalidListType": {
					"Title":  "Type de liste invalide",
					"Detail": "Le type de liste doit être 'film'.",
				},
				"MovieAlreadyInList": {
					"Title":  "Film déjà dans la liste",
					"Detail": "Le film est déjà présent dans la liste.",
				},
				"ErrorFetchingMovies": {
					"Title":  "Erreur lors de la récupération des films",
					"Detail": "Une erreur est survenue lors de la récupération des films avec les ID fournis.",
				},
				"ErrorAddingMovies": {
					"Title":  "Erreur lors de l'ajout des films",
					"Detail": "Une erreur est survenue lors de l'ajout des films à la liste.",
				},
			},
			"AuthMiddleware": {
				"UnauthorizedHeader": {
					"Title":  "En-tête d’Autorisation Manquant",
					"Detail": "L’en-tête d’autorisation est requis",
				},
				"UnauthorizedBearer": {
					"Title":  "Format d’Autorisation Invalide",
					"Detail": "L’en-tête d’autorisation doit être au format 'Bearer <token>'",
				},
				"UnauthorizedTokenParse": {
					"Title":  "Méthode de Signature Inattendue",
					"Detail": "Méthode de signature inattendue",
				},
				"UnauthorizedInvalidToken": {
					"Title":  "Jeton Invalide",
					"Detail": "Le jeton n'a pas pu être analysé ou est invalide",
				},
				"UnauthorizedToken": {
					"Title":  "Jeton Invalide",
					"Detail": "Le jeton n'est pas valide",
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
			"AddBrandsListUseCase": {
				"UserNotFound": {
					"Title":  "Usuario No Encontrado",
					"Detail": "No se pudo encontrar al usuario con el ID proporcionado.",
				},
				"UserNotActive": {
					"Title":  "Usuario No Activo",
					"Detail": "La cuenta del usuario no está activa. Por favor contacte con el soporte.",
				},
				"UserNotAdmin": {
					"Title":  "Usuario No es Administrador",
					"Detail": "El usuario no tiene privilegios de administrador.",
				},
				"ListNotFound": {
					"Title":  "Lista No Encontrada",
					"Detail": "No se pudo encontrar la lista con el ID proporcionado.",
				},
				"InvalidListType": {
					"Title":  "Tipo de Lista Inválido",
					"Detail": "El tipo de lista debe ser 'brand'.",
				},
				"BrandAlreadyInList": {
					"Title":  "Marca Ya En La Lista",
					"Detail": "La marca con el ID proporcionado ya existe en la lista.",
				},
				"ErrorFetchingBrands": {
					"Title":  "Error al Obtener Marcas",
					"Detail": "Ocurrió un error al obtener las marcas.",
				},
				"ErrorAddingBrands": {
					"Title":  "Error al Agregar Marcas",
					"Detail": "Ocurrió un error al agregar las marcas a la lista.",
				},
				"ErrorFetchingCombinations": {
					"Title":  "Error al Obtener Combinaciones",
					"Detail": "Ocurrió un error al obtener las combinaciones para la lista.",
				},
			},
			"AddMoviesListUseCase": {
				"UserNotFound": {
					"Title":  "Usuario no encontrado",
					"Detail": "No se pudo encontrar el usuario con el ID proporcionado.",
				},
				"UserNotActive": {
					"Title":  "Usuario no activo",
					"Detail": "La cuenta del usuario no está activa. Contacta con el soporte.",
				},
				"UserNotAdmin": {
					"Title":  "Acceso denegado",
					"Detail": "El usuario no tiene privilegios de administrador.",
				},
				"ListNotFound": {
					"Title":  "Lista no encontrada",
					"Detail": "No se pudo recuperar la lista con el ID proporcionado.",
				},
				"InvalidListType": {
					"Title":  "Tipo de lista no válido",
					"Detail": "El tipo de lista debe ser 'película'.",
				},
				"MovieAlreadyInList": {
					"Title":  "Película ya en la lista",
					"Detail": "La película ya está presente en la lista.",
				},
				"ErrorFetchingMovies": {
					"Title":  "Error al obtener películas",
					"Detail": "Ocurrió un error al obtener las películas con los IDs proporcionados.",
				},
				"ErrorAddingMovies": {
					"Title":  "Error al agregar películas",
					"Detail": "Ocurrió un error al agregar películas a la lista.",
				},
			},
			"AuthMiddleware": {
				"UnauthorizedHeader": {
					"Title":  "Falta el Encabezado de Autorización",
					"Detail": "El encabezado de autorización es obligatorio",
				},
				"UnauthorizedBearer": {
					"Title":  "Formato de Autorización Inválido",
					"Detail": "El encabezado de autorización debe estar en el formato 'Bearer <token>'",
				},
				"UnauthorizedTokenParse": {
					"Title":  "Método de Firma Inesperado",
					"Detail": "Método de firma inesperado",
				},
				"UnauthorizedInvalidToken": {
					"Title":  "Token Inválido",
					"Detail": "El token no pudo ser analizado o es inválido",
				},
				"UnauthorizedToken": {
					"Title":  "Token Inválido",
					"Detail": "El token no es válido",
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
			"AddBrandsListUseCase": {
				"UserNotFound": {
					"Title":  "用户未找到",
					"Detail": "无法找到提供的用户ID。",
				},
				"UserNotActive": {
					"Title":  "用户未激活",
					"Detail": "该用户账户未激活。请联系支持。",
				},
				"UserNotAdmin": {
					"Title":  "用户非管理员",
					"Detail": "用户没有管理员权限。",
				},
				"ListNotFound": {
					"Title":  "列表未找到",
					"Detail": "无法找到提供的列表ID。",
				},
				"InvalidListType": {
					"Title":  "无效的列表类型",
					"Detail": "列表类型必须为 'brand'。",
				},
				"BrandAlreadyInList": {
					"Title":  "品牌已在列表中",
					"Detail": "提供的品牌ID已经存在于列表中。",
				},
				"ErrorFetchingBrands": {
					"Title":  "获取品牌时出错",
					"Detail": "获取品牌时发生错误。",
				},
				"ErrorAddingBrands": {
					"Title":  "添加品牌时出错",
					"Detail": "添加品牌到列表时发生错误。",
				},
				"ErrorFetchingCombinations": {
					"Title":  "获取组合时出错",
					"Detail": "获取列表组合时发生错误。",
				},
			},
			"AddMoviesListUseCase": {
				"UserNotFound": {
					"Title":  "用户未找到",
					"Detail": "无法找到提供的用户ID。",
				},
				"UserNotActive": {
					"Title":  "用户未激活",
					"Detail": "用户账户未激活，请联系支持。",
				},
				"UserNotAdmin": {
					"Title":  "拒绝访问",
					"Detail": "用户没有管理员权限。",
				},
				"ListNotFound": {
					"Title":  "列表未找到",
					"Detail": "无法检索提供ID的列表。",
				},
				"InvalidListType": {
					"Title":  "无效的列表类型",
					"Detail": "列表类型必须为 '电影'。",
				},
				"MovieAlreadyInList": {
					"Title":  "电影已在列表中",
					"Detail": "的电影已在列表中。",
				},
				"ErrorFetchingMovies": {
					"Title":  "获取电影时出错",
					"Detail": "获取提供的电影ID时发生了错误。",
				},
				"ErrorAddingMovies": {
					"Title":  "添加电影时出错",
					"Detail": "将电影添加到列表时发生了错误。",
				},
			},
			"AuthMiddleware": {
				"UnauthorizedHeader": {
					"Title":  "缺少授权标头",
					"Detail": "授权标头是必需的",
				},
				"UnauthorizedBearer": {
					"Title":  "无效的授权格式",
					"Detail": "授权标头必须采用 'Bearer <token>' 格式",
				},
				"UnauthorizedTokenParse": {
					"Title":  "意外的签名方法",
					"Detail": "意外的签名方法",
				},
				"UnauthorizedInvalidToken": {
					"Title":  "无效的令牌",
					"Detail": "令牌无法解析或无效",
				},
				"UnauthorizedToken": {
					"Title":  "无效的令牌",
					"Detail": "令牌无效",
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

func GetErrorMessage(useCase, errorCode string) exceptions.ErrorMessage {
	if langMsgs, ok := errorMessages[currentLanguage]; ok {
		if uc, ok := langMsgs[useCase]; ok {
			if err, ok := uc[errorCode]; ok {
				return exceptions.ErrorMessage{
					Title:  err["Title"],
					Detail: err["Detail"],
				}
			}
		}
	}

	if uc, ok := errorMessages["en-US"][useCase]; ok {
		if err, ok := uc[errorCode]; ok {
			return exceptions.ErrorMessage{
				Title:  err["Title"],
				Detail: err["Detail"],
			}
		}
	}

	return exceptions.ErrorMessage{
		Title:  "Internal Error",
		Detail: "An unexpected error occurred.",
	}
}

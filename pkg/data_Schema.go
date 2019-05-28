package root

type User struct {
	Identifier string `json:"ID"`

	UserName string `json:"UserName"`
	Password string `json:"PSW"`
}

type UserService interface {
	CreateUser(u *User) error
	GetUserByUsername(username string) (User, error)
	Login(cred Credentials) (User, error, bool)
}

type Record struct {
	Identifier string `json:"ID"`

	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

type RecordService interface {
	CreateRecord(rec *Record) error
	GetAllRecords() ([]Record, error)
}

type ClaimDefn struct {
	UserIdentifier string `json:"UName"`
	CommonName     string `json:"CNAME"`

	ClaimDefnIdentifier string            `json:"CDID"`
	AttributesToType    map[string]string `json:"ATTR"`
}

type Claim struct {
	//	Identifier string `json:"Identifier"`

	UserName string `json:"UserName"`

	CommonName string `json:CommonName`
	IssuerName string `json:"IssuerName`

	//	PublicKey string `json:"PubKey"`

	Endpoint   string            `json:"Endpoint"`
	HashedData map[string]string `json:HashedData`
}

type ClaimService interface {
	CreateClaimDefn(map[string]string, string, string) error
	GetClaimDefnByCommonName(string, string) (ClaimDefn, error)
	GetAllClaimDefns() ([]ClaimDefn, error)

	CreateClaim(*Claim) error

	// GetClaimByUserID(string) ([]Claim, error)
	// GetClaimByCommonName(string, string) (Claim, error)
	// GetClaimDefnByClaimDefnID(string) ([]ClaimDefn, error)
	GetAllClaims() ([]Claim, error)
}

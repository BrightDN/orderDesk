# Error Patterns Analysis - Services Directory

## Overview
This document shows all files in the services directory that return errors and their current error structures.

---

## 1. AUTHENTICATION SERVICE
**Location:** `/internal/services/authentication/`

### verifyUser.go
**Error Types:**
- `ErrInvalidCredentials` - `errors.New("invalid credentials")`
- `ErrInternalError` - `errors.New("internal error")`

**Pattern:** Custom error variables defined at package level; returned directly

```go
func (auth *AuthenticationService) VerifyUser(c echo.Context, email, password string) (database.User, error) {
    // Returns: ErrInvalidCredentials, ErrInternalError
}
```

### signUp.go
**Error Types:**
- `errCreatingUser` - `errors.New("error: failed to create user")`
- `errCreatingCompanyEmployee` - `errors.New("error: failed to create company employee")`
- `errHashingPassword` - `errors.New("error: failed to hash password")`
- Raw `err` from database operations and `tx.Commit()`

**Pattern:** Mix of custom errors and raw database errors passed through

```go
func (auth *AuthenticationService) SignUp(c echo.Context, email, password, name string, invitation *database.Invite) (*database.CompanyUser, error) {
    // Returns: errHashingPassword, errCreatingUser, errCreatingCompanyEmployee, raw db errors
}
```

### hashPassword.go
**Error Types:**
- `fmt.Errorf("something went wrong: %v", err)`

**Pattern:** Formatted error wrapping with message

```go
func (auth *AuthenticationService) hashPassword(password string) (string, error) {
    // Returns: Formatted error string
}
```

### comparePasswordHash.go
**Error Types:**
- `*errorHandling.AppError` ✓ **ALREADY USING TARGET FORMAT**

**Pattern:** Returns structured AppError (the target format!)

```go
func (auth *AuthenticationService) comparePasswordHash(password string, hash string) (bool, *errorHandling.AppError) {
    // Returns: &errorHandling.AppError with Action and LogError fields
}
```

### logout.go
**Error Types:**
- `ErrSessionErr` - `errors.New("error: Failed clearing session")`
- Raw `err` from `session.Clear()`

**Pattern:** Custom error with flash messaging and raw errors

```go
func (auth *AuthenticationService) Logout(c echo.Context) error {
    // Returns: ErrSessionErr, raw session errors, flash errors
}
```

---

## 2. COMPANIES SERVICE
**Location:** `/internal/services/companies/`

### companies.go
**Error Types:**
- `ErrInternalError` - `errors.New("something went wrong on our end, try again later")`

**Pattern:** Custom error defined at package level

### create.go
**Error Types:**
- `ErrDuplicateName` - `errors.New("this name already exists")`
- `ErrDuplicateEmail` - `errors.New("this email already exists")`
- Wrapped with `fmt.Errorf("%w", ErrDuplicate*)`
- PostgreSQL error checking via `pq.Error`

**Pattern:** Custom errors with constraint checking and error wrapping

```go
func (cs *CompanyService) Create(c echo.Context) (database.Company, error) {
    // Returns: fmt.Errorf("%w", ErrDuplicateName/Email), or ErrInternalError
    // Uses pq.Error for constraint validation
}
```

### delete.go
**Error Types:**
- Raw database error

**Pattern:** Passes through raw database errors

```go
func (cs *CompanyService) Delete(c echo.Context, id int32) error {
    // Returns: raw database error
}
```

### getCompanies.go
**Error Types:**
- `ErrInternalError` (custom error)

**Pattern:** Returns generic custom error on any database failure

```go
func (cs *CompanyService) GetCompanies(c echo.Context) ([]Company, error) {
    // Returns: ErrInternalError
}
```

### getCompany.go
**Error Types:**
- `ErrInternalError` (custom error)

**Pattern:** Returns generic custom error on any database failure

```go
func (cs *CompanyService) GetCompany(c echo.Context, id int32) (Company, error) {
    // Returns: ErrInternalError
}
```

### getCompanyEmployees.go
**Error Types:**
- `ErrInternalError` (custom error)

**Pattern:** Returns generic custom error on any database failure

```go
func (cs *CompanyService) GetCompanyEmployees(c echo.Context, id int32) ([]Employee, error) {
    // Returns: ErrInternalError
}
```

### getEmployee.go
**Error Types:**
- Raw database error

**Pattern:** Passes through raw database error

```go
func (cs *CompanyService) GetEmployee(c echo.Context, userID int32) (Employee, error) {
    // Returns: raw database error
}
```

### update.go
**Error Types:**
- `ErrInternalError` (custom error)

**Pattern:** Returns generic custom error on any database failure

```go
func (cs *CompanyService) Update(c echo.Context, id int32) error {
    // Returns: ErrInternalError
}
```

### employee.go
**No errors** - Struct definition only

---

## 3. SUPPLIERS SERVICE (Subdirectory of Companies)
**Location:** `/internal/services/companies/suppliers/`

### suppliers.go
**Error Types:**
- `ErrInternalError` - `errors.New("Something went wrong on our end, please try again later")`

**Pattern:** Custom error defined at package level

### create.go
**Error Types:**
- `ErrSupplierNameExists` - `errors.New("a supplier with that name already exists")`
- `ErrSupplierEmailExists` - `errors.New("a supplier with that email already exists")`
- `ErrInternalError` (from suppliers.go)
- PostgreSQL error checking via `pq.Error`
- Uses `logging.ErrorLog()` for logging

**Pattern:** Custom errors with constraint checking, logging, and flash messaging

```go
func (s *SupplierService) Create(c echo.Context, company, email, contact string, companyID int32) error {
    // Returns: ErrSupplierNameExists, ErrSupplierEmailExists, ErrInternalError
    // Uses: logging.ErrorLog(), flash.Set() for error handling
    // Uses: pq.Error for constraint validation
}
```

### createProduct.go
**Error Types:**
- `ErrSupplierProductNotUnique` - `errors.New("Error: product already exists for supplier")`
- `ErrInternalError` (from suppliers.go)
- PostgreSQL error checking via `pq.Error`

**Pattern:** Custom error with constraint checking

```go
func (sc *SupplierService) CreateProduct(c echo.Context, supplierID int32, product string) error {
    // Returns: ErrSupplierProductNotUnique, ErrInternalError
    // Uses: pq.Error for constraint validation
}
```

### deleteProduct.go
**Error Types:**
- `ErrInternalError` (from suppliers.go)
- Uses `fmt.Printf()` for logging (not structured logging)

**Pattern:** Returns generic error; uses printf for logging

```go
func (sc *SupplierService) DeleteProduct(c echo.Context, supplierID, productID int32) error {
    // Returns: ErrInternalError
    // Uses: fmt.Printf() for logging (problematic)
}
```

### editInformation.go
**Error Types:**
- Raw database errors

**Pattern:** Passes through raw database errors

```go
func (ss *SupplierService) EditSupplier(c echo.Context, name string, compID int32, newName, email, contact, msubject, mctx string) (Supplier, error) {
    // Returns: raw database errors
}
```

### getAllByCompany.go
**Error Types:**
- Raw database error (after flash.Set attempt)
- Uses flash messaging on error

**Pattern:** Flash messaging with raw error passthrough

```go
func (s *SupplierService) GetAllByCompany(c echo.Context, companyID int32) ([]Supplier, error) {
    // Returns: raw database error + flash message attempt
}
```

### getProducts.go
**Error Types:**
- Raw database error (after flash.Set attempt)
- Uses flash messaging on error

**Pattern:** Flash messaging with raw error passthrough

```go
func (sc *SupplierService) GetProducts(c echo.Context, id int32) ([]Products, error) {
    // Returns: raw database error + flash message attempt
}
```

### getSupplier.go
**Error Types:**
- `ErrInternalError` (from suppliers.go)
- Uses `logging.ErrorLog()` in some cases

**Pattern:** Returns generic custom error; some logging

```go
func (sc *SupplierService) GetSupplierByID(c echo.Context, supplierID int32) (Supplier, error) {
    // Returns: ErrInternalError

func (sc *SupplierService) GetSupplierByNameAndCompanyID(c echo.Context, supplierName string, companyID int32) (Supplier, error) {
    // Returns: ErrInternalError
    // Uses: logging.ErrorLog()
}
```

---

## 4. INVITES SERVICE
**Location:** `/internal/services/invites/`

### invite.go
**Error Types:**
- `ErrMaxAttempts` - `errors.New("maximum attempts passed")`
- `ErrInviteCreation` - `errors.New("failed to generate an invitation")`
- `ErrTokenCreation` - `errors.New("failed to generate a token")`
- `ErrUnexpectedValue` - `errors.New("Unexpected value, action failed")`
- `ErrInternalError` - `errors.New("Something went wrong and we could not complete your request")`
- `ErrAlreadyAccepted` - `errors.New("The invite has already been accepted")`

**Pattern:** Custom error variables defined at package level

### generateToken.go
**Error Types:**
- `fmt.Errorf("%w: %v", ErrTokenCreation, err)`

**Pattern:** Error wrapping with custom error

```go
func (is *InvitationService) generateToken(length int) (string, error) {
    // Returns: Wrapped error with ErrTokenCreation
}
```

### getInvitation.go
**Error Types:**
- `ErrInvalidInvitationToken` - `errors.New("invalid invitation token")`
- `ErrExpiredToken` - `errors.New("the token has already expired")`
- `ErrAlreadyUsed` - `errors.New("this token has already been used")`
- `ErrInternalError` (from invite.go)

**Pattern:** Custom errors defined locally with specific logic

```go
func (is *InvitationService) GetInvitation(c echo.Context, token string) (Invitation, error) {
    // Returns: ErrInvalidInvitationToken, ErrExpiredToken, ErrAlreadyUsed, ErrInternalError
}
```

### getInvites.go
**Error Types:**
- Silently ignores errors: `cinvs, _ := is.db.GetCompanyInvites(...)`

**Pattern:** Error suppression with blank identifier (problematic!)

```go
func (is *InvitationService) GetCompanyInvites(c echo.Context) []Invite {
    // ISSUE: Silently ignores database errors
}
```

### delete.go
**Error Types:**
- Raw database error

**Pattern:** Passes through raw database error

```go
func (is *InvitationService) Delete(c echo.Context, id int32) error {
    // Returns: raw database error
}
```

### reactivate.go
**Error Types:**
- Raw database error
- Uses `fmt.Printf()` for logging
- Uses flash messaging

**Pattern:** Printf logging + flash + raw error

```go
func (is *InvitationService) Reactivate(c echo.Context, id int32) error {
    // Returns: raw database error + flash message attempt
    // Uses: fmt.Printf() for logging (problematic)
}
```

### resend.go
**Error Types:**
- Raw database error
- Uses flash messaging
- Raw error from mailer service

**Pattern:** Flash + raw errors from DB and mailer

```go
func (is *InvitationService) Resend(c echo.Context, id int32) error {
    // Returns: raw database error, raw mailer error + flash message attempts
}
```

### sendCompany.go
**Error Types:**
- Raw error from `companyService.Create()`
- Wrapped with `fmt.Errorf("%w: %v", ErrMaxAttempts/ErrInviteCreation, err)`
- `echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf(...))`

**Pattern:** Error wrapping, formatted errors, and echo HTTP errors

```go
func (is *InvitationService) SendCompany(c echo.Context) error {
    // Returns: Wrapped errors with ErrMaxAttempts/ErrInviteCreation
    // Returns: echo.NewHTTPError for mail sending failures
}
```

### validateInvitation.go
**Error Types:**
- `ErrInvalidInvitation` - `errors.New("invalid invitation")`
- `ErrInternalServerError` - `errors.New("internal server error")`
- Checks for `sql.ErrNoRows`

**Pattern:** Custom errors with specific validation

```go
func (inv *InvitationService) ValidateInvitation(c echo.Context, token, email string) (*database.Invite, error) {
    // Returns: ErrInvalidInvitation, ErrInternalServerError
}
```

### invites.go
**No error returns** - Struct definition only

### invitation.go
**No error returns** - Struct definition only

### getCompanyInvitationMail.go
**No error returns** - String formatting only

---

## 5. MAILER SERVICE
**Location:** `/internal/services/mailer/`

### mailer.go
**Error Types:**
- Raw error from `NewClient()`

**Pattern:** Passes through raw initialization error

```go
func NewMailerService(cfg configs.MailConfig) (*MailerService, error) {
    // Returns: raw error from NewClient
}
```

### send.go
**Error Types:**
- `ErrMailSending` - `errors.New("failed to send mail")`
- Wrapped with `fmt.Errorf("%w: %v", ErrMailSending, err)`

**Pattern:** Custom error wrapped in formatted error

```go
func (ms *MailerService) Send(m Mail) error {
    // Returns: Wrapped error with ErrMailSending
}
```

### client.go
**Error Types:**
- `fmt.Errorf("Mail provider can not be empty")`
- `fmt.Errorf("username can not be empty")`
- `fmt.Errorf("password can not be empty")`
- `fmt.Errorf("Something went wrong initializing the mailing provider: %v", err)`

**Pattern:** Formatted validation errors

```go
func NewClient(cfg configs.MailConfig) (*mail.Client, error) {
    // Returns: fmt.Errorf(...) for validation failures
}
```

### close.go
**Error Types:**
- Raw error from underlying client

**Pattern:** Passes through raw error

```go
func (m *MailerService) Close() error {
    // Returns: raw error from m.client.Close()
}
```

### mail.go
**No errors** - Struct definition only

---

## 6. PERMISSIONS SERVICE
**Location:** `/internal/services/permissions/`

### permissions.go, admin.go, employee.go, custom.go
**No error returns** - All methods return permission structs, no error handling

---

## SUMMARY OF ERROR PATTERNS

### Pattern 1: Custom Errors with `errors.New()`
Most common pattern. Custom error variables defined and returned directly.

**Files:** 
- verifyUser.go, logout.go
- companies.go, create.go
- suppliers.go, create.go, createProduct.go
- invite.go, getInvitation.go, validateInvitation.go
- mailer send.go

### Pattern 2: Formatted Errors with `fmt.Errorf()`
Errors formatted with messages and context.

**Files:**
- hashPassword.go
- create.go (with `%w` for wrapping)
- generateToken.go (with `%w` for wrapping)
- sendCompany.go (with `%w` for wrapping)
- client.go
- send.go (with `%w` for wrapping)

### Pattern 3: Raw Error Passthrough
Errors from database or other services passed directly.

**Files:**
- signUp.go
- delete.go
- getEmployee.go
- editInformation.go
- getAllByCompany.go, getProducts.go
- delete.go (invites)
- reactivate.go, resend.go
- close.go

### Pattern 4: Echo HTTP Errors
Using `echo.NewHTTPError()` for HTTP responses.

**Files:**
- sendCompany.go

### Pattern 5: PostgreSQL Constraint Checking
Using `pq.Error` to validate unique constraints and return specific errors.

**Files:**
- create.go (companies)
- create.go, createProduct.go (suppliers)
- sendCompany.go (invites)

### Pattern 6: Flash Messaging
Using flash message system alongside error returns.

**Files:**
- logout.go
- create.go, Reactivate.go, resend.go (suppliers/invites)
- getAllByCompany.go, getProducts.go

### Pattern 7: Logging Integration
Using custom logging functions.

**Files:**
- create.go (suppliers)
- getSupplier.go (suppliers)

### Pattern 8: AppError (TARGET FORMAT) ✓
Already using the target structured error format.

**Files:**
- comparePasswordHash.go

### Pattern 9: Error Suppression (PROBLEMATIC)
Ignoring errors with blank identifier.

**Files:**
- getInvites.go - `cinvs, _ := is.db.GetCompanyInvites(...)`

### Pattern 10: Printf Logging (PROBLEMATIC)
Using `fmt.Printf()` instead of structured logging.

**Files:**
- deleteProduct.go
- reactivate.go

---

## CONVERSION RECOMMENDATIONS

### High Priority (Many errors, inconsistent handling)
1. **invites/sendCompany.go** - Mix of fmt.Errorf, echo.NewHTTPError, and raw errors
2. **suppliers/create.go** - Complex error handling with logging and flash
3. **authentication/signUp.go** - Mix of custom and raw errors

### Medium Priority (Inconsistent patterns)
4. **companies/create.go** - Constraint checking with wrapped errors
5. **suppliers/** (multiple files) - Flash messaging mixed with errors
6. **invites/resend.go, reactivate.go** - Printf logging issues

### Low Priority (Already good or few errors)
7. **comparePasswordHash.go** - Already using AppError ✓
8. **permissions/** - No error handling needed
9. **mailer/mail.go, invites/invitation.go, invites/invite.go** - Struct definitions

---

## INCONSISTENCIES FOUND

1. **Error message formatting**: Some use "error: " prefix, some don't, some use full sentences
2. **Logging approach**: Mix of fmt.Printf, logging.ErrorLog, and flash.Set
3. **Raw error passthrough**: No consistent wrapping with context
4. **Constraint checking**: Only done in some files (create.go)
5. **HTTP errors**: Some use echo.NewHTTPError, others return plain errors
6. **Error suppression**: getInvites.go silently ignores errors

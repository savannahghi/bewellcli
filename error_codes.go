package base

// ErrorCode are  used to determine the nature of an error, and why it occurred
// both the frontend and backend should be aware of these codes
type ErrorCode int

const (
	// OK is returned on success.
	OK ErrorCode = iota + 1

	// Internal errors means some invariants expected by underlying
	// system has been broken. If you see one of these errors,
	// something is very broken.
	// it's value is 2
	Internal

	// UndefinedArguments errors means either one or more arguments to
	// a method have not been specified
	// it's value is 3
	UndefinedArguments

	// PhoneNumberInUse indicates that a phone number has an associated user profile.
	// this error can occur when fetching a user profile using a phone number, to check
	// that the phone number has not already been registered. The check usually runs
	// on both PRIMARY PHONE and SECONDARY PHONE
	// it's value is 4
	PhoneNumberInUse

	// EmailAddressInUse indicates that an email address has an associated user profile.
	// this error can occur when fetching a user profile using an email address, to check
	// that the email address has not already been registered. The check usually runs
	// on both PRIMARY EMAIL ADDRESS and SECONDARY EMAIL ADDRESS.
	// it's value is 5
	EmailAddressInUse

	// UsernameInUse indicates that a username has an associated user profile.
	// this error can occur when trying a update a user's username with a username that already has been taken
	// it's value is 6
	UsernameInUse

	// ProfileNotFound errors means a user profile does not exist with the provided parameters
	// This occures when fetching a user profile either by UID, ID , PHONE NUMBER or EMAIL and no
	// matching record is found
	// it's value is 7
	ProfileNotFound

	// PINMismatch errors means that the provided PINS do not match (are not similar)
	// it's value is 8
	PINMismatch

	// PINNotFound errors means a user PIN does not exist with the provided parameters
	// This occurs when fetching a PIN by the user's user profile ID and no
	// matching record is found. This should never occur and if it does then it means
	// there is a serious issue with our data
	// it's value is 9
	PINNotFound

	// UserNotFound errors means that a user's firebase auth account does not exists. This occurs
	// when fetching a firebase user by either a phone number or an email and their record is not found
	// it's value is 10
	UserNotFound

	// ProfileSuspended error means that user's profile has been suspended.
	// This may occur due to violation of terms or detection of suspicious activity
	// It's value is 11
	ProfileSuspended

	// PINError error means that some actions could not be performed on the PIN.
	// This may occur when the provided PIN cannot be encrypted, cannot be validated and/or is of invalid length
	// It's value is 12
	PINError

	// InvalidPushTokenLength means that an invalid push token was given.
	// This may occur when the lenth of the issued token is of less then the minimum character(1250)
	// It's error code is 13
	InvalidPushTokenLength

	// InvalidEnum means that the provided enumerator was of invalid.
	// This may occur when an invalid enum value has been defined. For example, PartnerType, LoginProviderType e.t.c
	// It's error code is 14
	InvalidEnum

	// OTPVerificationFailed means that the provide OTP could not be verified
	// This may occur when an incorrect OTP is supplied
	// It's error code is 15
	OTPVerificationFailed

	// MissingInput means that no OTP was submiited
	// This may occur when a user fails to provide an OTP but makes a submission
	// It's error code id 16
	MissingInput

	// InvalidFlavour means that the provide falvour is invalid
	// This may happen when the provided flavour is not consumer or pro
	// It's error code is 17
	InvalidFlavour

	// RecordNotFound means that the provided record is not found.
	// This may happen when the provided data e.g currency, user etc is not accepted
	// It's error code is 19
	RecordNotFound

	// UnableToFindProvider means that the selected provider could not be found
	// This may happen if the provider is not specified in the charge master
	// It's error code is 20
	UnableToFindProvider

	// PublishNudgeFailure means that there was an error while publishing a nudge
	// It's error code is 21
	PublishNudgeFailure

	// InvalidCredentials means that the provided credentials are invalid
	// This may happen when any of the customers provides wrong credentials
	// It's error code is 22
	InvalidCredentials

	// AddNewRecordError means that the record could not be saved
	// This may happen may be as a result of wrong credentials or biodata
	// It's error code is 23
	AddNewRecordError

	// InvalidSladeCode means that there is a problem with the provide slade code
	// This may occur when the slade code in the EDI does not match with the chosen provider
	// It's error code is 24
	InvalidSladeCode

	// KYCAlreadySubmitted means that there is a problem while submitting KYC
	// This may happen when the KYC has already been subnmitted
	// Its error code is 25
	KYCAlreadySubmitted

	// RoleNotValid means that the user role does not match the role required
	// to perform the current operation that the user is trying to perform.
	// Its error code is 26
	RoleNotValid

	//InsuranceDetailsNotCoverLinked means that the user's insurance details are not
	//coverLinked
	//Its error code is 27
	InsuranceDetailsNotCoverLinked

	//CoverAlreadyExists means that the user has tried to add an already existing cover
	//Its error code is 28
	CoverAlreadyExists

	//UserNotAuthorizedToAccessThisResource means that the subject's
	//email has been found to not have access to the specified resource
	//Its error code is 29
	UserNotAuthorizedToAccessThisResource

	//NilEligibilityMember means that the eligibility lookup
	//done on the Slade360 EDI returns a nil
	//Its error code is 30
	NilEligibilityMember

	//MinorMember means that
	//a minor has tried to add their own cover
	//Its error code is 31
	MinorMember

	// NilPayerDetails means that
	//the payer details have been found to be nil
	//Its error code is 32
	NilPayerDetails

	// FailedToLookUpEligibility means that
	//there is a failed eligilibility lookup on the connected Slade 360EDI
	//This lookup also includes dependants and allowed providers in the response payload
	//Its error code is 33
	FailedToLookUpEligibility

	//UnlinkedCover means that
	//a subject's cover details have been found to not be linked
	//Its error code is 34
	UnlinkedCover

	//NoBeneficiaryWithMemberNumber means that
	//no beneficiaryId is found under the given member number
	//Its error code is 35
	NoBeneficiaryWithMemberNumber

	//FailedToFetchDependants means that
	//the dependants of a particular principal cannot be retrieved
	//Its error code is 36
	FailedToFetchDependants

	//InvalidCover means that
	//the eligibilityStatus does not equal a valid cover
	//Its error code is 37
	InvalidCover

	//UnableToAddCover means that
	//the cover has not been added
	//Its error code is 38
	UnableToAddCover

	//NotARecognizedDocumentType means that
	//the document used is not a recognized identification document
	//Its error code is 39
	NotARecognizedDocumentType

	// FailedCreateMemberProfile means that
	//when the request to create a new member profile fails
	// Its error code is 40
	FailedCreateMemberProfile

	//FailedCreateMemberProfileNotificationPayload is the  error message displayed
	//when creation of a member profile notification payload fails
	// Its error code is 41
	FailedCreateMemberProfileNotificationPayload

	//UnableToFindUserBioData means that
	//a subject's bio data cannot be found
	//Its error code is 42
	UnableToFindUserBioData

	//UnableToSaveMemberProfileRequest means that
	//saving a member's profile request has failed
	//Its error code is 43
	UnableToSaveMemberProfileRequest

	//UnableToCheckIfUserIsAnAdmin means that
	//checking to see if a user is an admin has failed
	//Its error code is 44
	UnableToCheckIfUserIsAnAdmin

	//LoggedInUserIsNotAnAdmin means that
	//the user currently logged in has been found to not be an admin
	//Its error code is 45
	LoggedInUserIsNotAnAdmin

	//UnableToRetrieveMemberRequestNotifications means that
	//retrieving member request notifications has failed
	//Its error code is 46
	UnableToRetrieveMemberRequestNotifications

	//UnableToMapMemberRequestNotification means that
	//mapping the member request notification doc to object has failed
	//Its error code is 47
	UnableToMapMemberRequestNotification

	// UnableToRetrieveNotification means that
	//retrieving a node from firestore fails with this ID
	//Its error code is 48
	UnableToRetrieveNotification

	//WrongNodeType means that
	//a programming error has been made and a notification payload was expected
	//Its error code is 49
	WrongNodeType

	//NilNotificationPayloadNode means that
	//a nil notification payload node is returned
	//Its error code is 50
	NilNotificationPayloadNode

	//UnableToSaveNotification means that
	//saving a notification after updating it to read has failed
	//Its error code is 51
	UnableToSaveNotification

	//NoConfirmedPhoneNumbers means that
	//a user's primary phone number is  nil
	//Its error code is 52
	NoConfirmedPhoneNumbers

	//InvalidPhoneNumberFormat means that
	//the phone number format is invalid
	//Its error code is 53
	InvalidPhoneNumberFormat

	// UnableToSendText means that
	//sending a text to the phone number in question has failed
	//Its  error code is 54
	UnableToSendText

	//UnknownStateProvided means that
	//an unknown state has been entered
	//Its error code is 55
	UnknownStateProvided

	//UnableToFetchProviderList means that
	//fetching a provider list has failed
	//Its error code is 56
	UnableToFetchProviderList

	//ExpectedCardToBeTypeString means that
	//the card is of another type apart from string
	//Its error code is 57
	ExpectedCardToBeTypeString

	//HasSladeCard means that
	//the type of HasSladeCard variable is of another type other than bool
	//Its error code is 58
	HasSladeCard
)

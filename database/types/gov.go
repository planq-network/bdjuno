package types

import "time"

// ProposalRow represents a single row inside the proposal table
type ProposalRow struct {
	Title           string    `db:"title"`
	Description     string    `db:"description"`
	ProposalRoute   string    `db:"proposal_route"`
	ProposalType    string    `db:"proposal_type"`
	ProposalID      uint64    `db:"proposal_id"`
	SubmitTime      time.Time `db:"submit_time"`
	DepositEndTime  time.Time `db:"deposit_end_time"`
	VotingStartTime time.Time `db:"voting_start_time"`
	VotingEndTime   time.Time `db:"voting_end_time" `
	Proposer        string    `db:"proposer"`
	Status          string    `db:"status"`
}

// NewProposalRow allows to easily create a new ProposalRow
func NewProposalRow(title string,
	description string,
	proposalRoute string,
	proposalType string,
	proposalID uint64,
	submitTime time.Time,
	depositEndTime time.Time,
	votingStartTime time.Time,
	votingEndTime time.Time,
	proposer string,
	status string) ProposalRow {
	return ProposalRow{
		Title:           title,
		Description:     description,
		ProposalRoute:   proposalRoute,
		ProposalType:    proposalType,
		ProposalID:      proposalID,
		SubmitTime:      submitTime,
		DepositEndTime:  depositEndTime,
		VotingStartTime: votingStartTime,
		VotingEndTime:   votingEndTime,
		Proposer:        proposer,
		Status:          status,
	}
}

// Equals return true if two ProposalRow are the same
func (w ProposalRow) Equals(v ProposalRow) bool {
	return w.Title == v.Title &&
		w.Description == v.Description &&
		w.ProposalRoute == v.ProposalRoute &&
		w.ProposalType == v.ProposalType &&
		w.ProposalID == v.ProposalID &&
		w.SubmitTime.Equal(v.SubmitTime) &&
		w.DepositEndTime.Equal(v.DepositEndTime) &&
		w.VotingStartTime.Equal(v.VotingStartTime) &&
		w.VotingEndTime.Equal(v.VotingEndTime) &&
		w.Proposer == v.Proposer &&
		w.Status == v.Status
}

// TallyResultRow represents a single row inside the tally_result table
type TallyResultRow struct {
	ProposalID int64 `db:"proposal_id"`
	Yes        int64 `db:"yes"`
	Abstain    int64 `db:"abstain"`
	No         int64 `db:"no"`
	NoWithVeto int64 `db:"no_with_veto"`
	Height     int64 `db:"height"`
}

// NewTallyResultRow return a new TallyResultRow instance
func NewTallyResultRow(
	proposalID int64,
	yes int64,
	abstain int64,
	no int64,
	noWithVeto int64,
	height int64,
) TallyResultRow {
	return TallyResultRow{
		ProposalID: proposalID,
		Yes:        yes,
		Abstain:    abstain,
		No:         no,
		NoWithVeto: noWithVeto,
		Height:     height,
	}
}

// Equals return true if two TallyResultRow are the same
func (w TallyResultRow) Equals(v TallyResultRow) bool {
	return w.ProposalID == v.ProposalID &&
		w.Yes == v.Yes &&
		w.Abstain == v.Abstain &&
		w.No == v.No &&
		w.NoWithVeto == v.NoWithVeto &&
		w.Height == v.Height
}

// VoteRow represents a single row inside the vote table
type VoteRow struct {
	ProposalID int64  `db:"proposal_id"`
	Voter      string `db:"voter"`
	Option     string `db:"option"`
	Height     int64  `db:"height"`
}

// NewVoteRow allows to easily create a new VoteRow
func NewVoteRow(
	proposalID int64,
	voter string,
	option string,
	height int64,
) VoteRow {
	return VoteRow{
		ProposalID: proposalID,
		Voter:      voter,
		Option:     option,
		Height:     height,
	}
}

// Equals return true if two VoteRow are the same
func (w VoteRow) Equals(v VoteRow) bool {
	return w.ProposalID == v.ProposalID &&
		w.Voter == v.Voter &&
		w.Option == v.Option &&
		w.Height == v.Height
}

// DepositRow represents a single row inside the deposit table
type DepositRow struct {
	ProposalID   int64   `db:"proposal_id"`
	Depositor    string  `db:"depositor"`
	Amount       DbCoins `db:"amount"`
	TotalDeposit DbCoins `db:"total_deposit"`
	Height       int64   `db:"height"`
}

// NewDepositRow allows to easily create a new NewDepositRow
func NewDepositRow(
	proposalID int64,
	depositor string,
	amount DbCoins,
	totalDeposit DbCoins,
	height int64,
) DepositRow {
	return DepositRow{
		ProposalID:   proposalID,
		Depositor:    depositor,
		Amount:       amount,
		TotalDeposit: totalDeposit,
		Height:       height,
	}
}

// Equals return true if two VoteDepositRow are the same
func (w DepositRow) Equals(v DepositRow) bool {
	return w.ProposalID == v.ProposalID &&
		w.Depositor == v.Depositor &&
		w.Amount.Equal(&v.Amount) &&
		w.TotalDeposit.Equal(&v.TotalDeposit) &&
		w.Height == v.Height
}

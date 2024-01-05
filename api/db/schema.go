package db

var (
	//nolint:unused
	schema = `
CREATE TABLE Challenges (
    AssertionHash TEXT NOT NULL PRIMARY KEY,
    UNIQUE(AssertionHash)
);

CREATE TABLE Edges (
    Id TEXT NOT NULL PRIMARY KEY,
    ChallengeLevel INTEGER NOT NULL,
    OriginId TEXT NOT NULL,
    StartHistoryRoot TEXT NOT NULL,
    StartHeight INTEGER NOT NULL,
    EndHistoryRoot TEXT NOT NULL,
    EndHeight INTEGER NOT NULL,
    CreatedAtBlock INTEGER NOT NULL,
    MutualId TEXT NOT NULL,
    ClaimId TEXT,
    HasChildren BOOLEAN NOT NULL,
    LowerChildId TEXT NOT NULL,
    UpperChildId TEXT NOT NULL,
    MiniStaker TEXT,
    AssertionHash TEXT NOT NULL,
    HasRival BOOLEAN NOT NULL,
    Status TEXT NOT NULL,
    HasLengthOneRival BOOLEAN NOT NULL,
    FOREIGN KEY(LowerChildID) REFERENCES Edges(Id),
    FOREIGN KEY(UpperChildID) REFERENCES Edges(Id),
    FOREIGN KEY(AssertionHash) REFERENCES Challenges(AssertionHash)
);

CREATE TABLE Assertions (
    Hash TEXT NOT NULL PRIMARY KEY,
    ConfirmPeriodBlocks INTEGER NOT NULL,
    RequiredStake TEXT NOT NULL,
    ParentAssertionHash TEXT NOT NULL,
    InboxMaxCount TEXT NOT NULL,
    AfterInboxBatchAcc TEXT NOT NULL,
    WasmModuleRoot TEXT NOT NULL,
    ChallengeManager TEXT NOT NULL,
    CreationBlock INTEGER NOT NULL,
    TransactionHash TEXT NOT NULL,
    BeforeStateBlockHash TEXT NOT NULL,
    BeforeStateSendRoot TEXT NOT NULL,
    BeforeStateMachineStatus TEXT NOT NULL,
    AfterStateBlockHash TEXT NOT NULL,
    AfterStateSendRoot TEXT NOT NULL,
    AfterStateMachineStatus TEXT NOT NULL,
    FirstChildBlock INTEGER,
    SecondChildBlock INTEGER,
    IsFirstChild BOOLEAN NOT NULL,
    Status TEXT NOT NULL,
    ConfigHash TEXT NOT NULL,
    FOREIGN KEY(Hash) REFERENCES Challenges(AssertionHash),
    FOREIGN KEY(ParentAssertionHash) REFERENCES Assertions(Hash)
);

CREATE INDEX idx_edge_assertion ON Edges(AssertionHash);
CREATE INDEX idx_assertions_assertion ON Assertions(Hash);
CREATE INDEX idx_edge_claim_id ON Edges(ClaimId);
CREATE INDEX idx_edge_end_height ON Edges(EndHeight);
CREATE INDEX idx_edge_end_history_root ON Edges(EndHistoryRoot);
`
)

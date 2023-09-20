create table Patients(
    `Id` int not null primary key auto_increment,
    `Name` text not null,
    `LastName` text not null,
    `Address` text not null,
    `DNI` text not null,
    `DischargeDate` text not null    

);
create table Dentists(
    `Id` int not null primary key auto_increment,
    `Name` text not null,
    `LastName` text not null,
    `Registration` text not null   
);
create table Turns(
    `Id` int not null primary key auto_increment,
    `PatientId` int not null,
    `DentistId` int not null,
    FOREIGN KEY (DentistId) REFERENCES Dentists(id),
    FOREIGN KEY (PatientId) REFERENCES Patients(id),
    `DateHour` text null,
    `Description` text null
);

INSERT INTO Patients (Name, LastName, Address, DNI, DischargeDate)
VALUES ('Tom', 'Erichsen', 'Skagen 21', '123456', '01-09-2023');

INSERT INTO Dentists (Name, LastName, Registration)
VALUES ('Joe', 'Perez', 'C123456');

INSERT INTO Turns (PatientId, DentistId, DateHour, Description)
VALUES (1, 1, '07-10-2023 10:00', 'lorem ipsum');
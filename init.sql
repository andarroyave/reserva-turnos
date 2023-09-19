-- Create the Dentists table
CREATE TABLE Dentists (
    Matriculation INT AUTO_INCREMENT PRIMARY KEY,
    Lastname VARCHAR(255) NOT NULL,
    Firstname VARCHAR(255) NOT NULL
);

-- Create the Patients table
CREATE TABLE Patients (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Lastname VARCHAR(255) NOT NULL,
    Firstname VARCHAR(255) NOT NULL,
    Address VARCHAR(255),
    RegistrationDate DATE
);

-- Create the Appointments table
CREATE TABLE Appointments (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Description VARCHAR(255),
    DateTime DATETIME NOT NULL,
    DentistMatriculation INT,
    PatientID INT,
    FOREIGN KEY (DentistMatriculation) REFERENCES Dentists(Matriculation),
    FOREIGN KEY (PatientID) REFERENCES Patients(ID)
);

-- Unique index on the Appointments table to prevent duplicates on the same date and time.
CREATE UNIQUE INDEX idx_appointments_date_time ON Appointments (DateTime);

-- End of SQL file

# Smart Emergency Response System (SERS)

A full-stack emergency response application that enables users to quickly request help during emergencies while automatically sharing their location and medical information with responders.

## Features

- 🆘 One-tap SOS activation with real-time location sharing
- 📍 Automatic nearby responder matching
- 📱 SMS alerts to emergency contacts via Twilio
- 🏥 Medical information sharing through QR codes
- 🔐 Secure authentication with JWT
- 📊 Health monitoring integration capability

## Tech Stack

### Backend

- **Language:** Go 1.19
- **Framework:** Gin
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT
- **SMS Service:** Twilio

### Frontend

- **Framework:** Flutter
- **State Management:** Provider
- **Location Services:** Geolocator
- **HTTP Client:** Dart HTTP

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- Flutter SDK
- PostgreSQL
- Twilio Account (for SMS features)

### Environment Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/sers.git
cd sers
```

2. Create backend environment file:

```bash
cd backend
cp .env.example .env
```

3. Update `.env` with your credentials:

```env
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=sers
JWT_SECRET=your_secret_key
TWILIO_ACCOUNT_SID=your_twilio_sid
TWILIO_AUTH_TOKEN=your_twilio_token
TWILIO_PHONE_NUMBER=your_twilio_number
```

### Running the Application

1. Start the backend services:

```bash
docker-compose up --build
```

2. Run the Flutter application:

```bash
cd frontend
flutter pub get
flutter run
```

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - User login

### Emergency

- `POST /api/sos/trigger` - Trigger SOS alert

## Database Schema

### Users

- ID
- Email
- Password (hashed)
- FullName
- Phone
- BloodGroup
- Allergies

### Emergency Contacts

- ID
- UserID
- Name
- Phone
- Relationship

### Locations

- ID
- UserID
- Latitude
- Longitude
- UpdatedAt

### Medical Records

- ID
- UserID
- RecordType
- Details
- Timestamp

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Twilio for SMS services
- Google Maps for location services
- Flutter team for the amazing framework
- Go team for the powerful backend language

# snapp sms bomber :LOL


# Armineyvazi

Armineyvazi is a command-line application written in Go that sends OTP (One-Time Password) requests to the Snapp Taxi API.

## Usage

1. Make sure you have Go installed on your system.

2. Clone the repository or download the code files.

3. Navigate to the project directory in your terminal.

4. Build the Go application by running the following command:

## go build

5. Run the application, providing the phone number as a command-line argument. For example:


Replace `<phone_number>` with the desired phone number to which you want to send the OTP requests.

**Note**: The phone number should be entered without any special characters or country codes. For example, if the phone number is +98 1234567890, you should enter only 1234567890 as the command-line argument.

6. The application will continuously send OTP requests to the Snapp Taxi API every 2 seconds. You can stop the application by pressing Ctrl+C in the terminal.

## Dependencies

The application relies on the following dependencies:

- `fmt`: Provides basic input/output functionality.
- `net/http`: Enables making HTTP requests.
- `net/url`: Provides URL encoding functionality.
- `os`: Provides access to command-line arguments.
- `time`: Allows adding delays between requests.

## License

This project is licensed under the [MIT License](LICENSE).

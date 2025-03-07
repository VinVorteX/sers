import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/user.dart';

class ApiService {
  final String baseUrl = 'http://localhost:8080/api';
  String? _token;

  void setToken(String token) {
    _token = token;
  }

  Future<Map<String, dynamic>> login(String email, String password) async {
    final response = await http.post(
      Uri.parse('$baseUrl/auth/login'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'email': email,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      _token = data['token'];
      return data;
    } else {
      throw Exception('Failed to login');
    }
  }

  Future<void> triggerSOS(double latitude, double longitude, String message) async {
    if (_token == null) throw Exception('Not authenticated');

    final response = await http.post(
      Uri.parse('$baseUrl/sos/trigger'),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $_token',
      },
      body: json.encode({
        'latitude': latitude,
        'longitude': longitude,
        'message': message,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to trigger SOS');
    }
  }
} 
import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import '../services/api_service.dart';

class SOSButton extends StatelessWidget {
  final ApiService _apiService = ApiService();

  Future<void> _triggerSOS(BuildContext context) async {
    try {
      // Get current location
      Position position = await Geolocator.getCurrentPosition(
        desiredAccuracy: LocationAccuracy.high
      );

      // Trigger SOS
      await _apiService.triggerSOS(
        position.latitude,
        position.longitude,
        "Emergency assistance needed!"
      );

      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Emergency services have been notified'))
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Failed to trigger SOS: ${e.toString()}'))
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => _triggerSOS(context),
      child: Container(
        width: 200,
        height: 200,
        decoration: BoxDecoration(
          shape: BoxShape.circle,
          color: Colors.red,
          boxShadow: [
            BoxShadow(
              color: Colors.black26,
              blurRadius: 10,
              spreadRadius: 5,
            )
          ],
        ),
        child: Center(
          child: Text(
            'SOS',
            style: TextStyle(
              color: Colors.white,
              fontSize: 48,
              fontWeight: FontWeight.bold,
            ),
          ),
        ),
      ),
    );
  }
} 
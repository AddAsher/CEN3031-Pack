import { Injectable } from '@angular/core';
import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent {
  rmsg: string;
  username: string;
  password: string;

  constructor(private authService: AuthService, private router: Router) {
    this.rmsg = "Register here";
    this.username = "";
    this.password = "";
  }


  onRegister() {
    console.log("Submitting register request:", this.username, this.password);
    this.authService.register(this.username, this.password).subscribe(
      (response) => {
        console.log("Request request received on Front end");
        this.goToLogin();
        document.forms[0].reset()
      },
      (error) => {
        document.forms[0].reset()
        console.log(error);
        const errorMessage = error.error.split('\n')[0];
        alert(errorMessage);
      }
    );
  }
  goToLogin() {
    this.router.navigate(['']);
  }
}

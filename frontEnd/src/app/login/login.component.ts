import { Injectable } from '@angular/core';
import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  password: string = "";
  username: string = "";
  title = 'PACK';
  users: any[] = [];

  constructor(private authService: AuthService, private router: Router) { }

  onSubmit() {
    console.log("Submitting login request:", this.username, this.password);
    this.authService.login(this.username, this.password).subscribe(
      (response) => {
        console.log("Login request received on Front end");
        this.goToHomePage();
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

  goToHomePage() {
    this.router.navigate(['/home']);
  }

  goToRegisterPage() {
    this.router.navigate(['/registration']);
  }

}

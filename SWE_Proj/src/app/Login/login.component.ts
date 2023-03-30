import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';

export interface User {
  id: number;
  username: string;
  email: string;
  password: string;
}

@Injectable({
  providedIn: 'root'
})




@Component({
  selector: 'app-root',
  template: `
  // <ul>
  // <li *ngFor="let user of users">{{ user.name }}</li>
  // </ul>
  // `,
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
        this.onLogin();
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

  onLogin(): void {
    // Do some login logic here...

    // Navigate to the home page
    this.router.navigate(['/home']);
  }

  onRegister(): void {
    // Do some login logic here...

    // Navigate to the home page
    this.router.navigate(['/register']);
  }


}

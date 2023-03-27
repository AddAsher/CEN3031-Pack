import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';


export interface User {
  id: number;
  username: string;
  email: string;
  password: string;
}

@Injectable({
  providedIn: 'root'
})



export class AuthService {

  private baseUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  login(username: string, password: string) {
    const url = `${this.baseUrl}/login`;
    const body = { username, password };
    return this.http.post<User>(url, body);
  }

}

@Component({
  selector: 'app-root',
  template: `
  <ul>
      <li *ngFor="let user of users">{{ user.name }}</li>
  </ul>
`,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent {
  email: string = "";
  password: string = "";
  username: string = "";
  title = 'PACK';
  users: any[] = [];

  constructor(private http: HttpClient) { }

  onSubmit() {
    this.http.post('http://localhost:8080/login', {
      email: this.email,
      password: this.password
    }).subscribe((response) => {
      console.log("Login request recieved");
    }, (error) => {
      console.log(error);
    });
  }

  buttonPress() {
    alert("Login button clicked")
    document.forms[0].reset();
  }
}

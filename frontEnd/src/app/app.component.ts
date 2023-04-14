import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  // password: string = "";
  // username: string = "";
  // title = 'PACK';
  // users: any[] = [];

  // constructor(private authService: AuthService, private router: Router) { }

  // onSubmit() {
  //   console.log("Submitting login request:", this.username, this.password);
  //   this.authService.login(this.username, this.password).subscribe(
  //     (response) => {
  //       console.log("Login request received on Front end");
  //       this.goToHomePage();
  //       document.forms[0].reset()
  //     },
  //     (error) => {
  //       document.forms[0].reset()
  //       console.log(error);
  //       const errorMessage = error.error.split('\n')[0];
  //       alert(errorMessage);
  //     }
  //   );
  // }

  // goToHomePage(){
  //   // Do some login logic here...


  //   // Navigate to the home page
  //   this.router.navigate(['/home']);
  // }

  // goToRegisterPage(){
  //   // Do some login logic here...

  //   // Navigate to the home page
  //   this.router.navigate(['/registration']);
  // }

}

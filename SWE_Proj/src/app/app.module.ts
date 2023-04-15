import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AuthService } from './Login/auth.service';
import { RouterModule, Routes } from '@angular/router';


import { LoginComponent } from './Login/login.component'; //login page
import { HomeComponent } from './Home/home.component'; //home page 
import { RegistrationComponent } from './Registration/registration.component'; //registration page
import { AppRoutingModule } from './app-routing.module';


@NgModule({
  declarations: [LoginComponent,HomeComponent,RegistrationComponent],
  imports: [BrowserModule, HttpClientModule, FormsModule, AppRoutingModule],
  providers: [AuthService],
  bootstrap: [LoginComponent],
  
})
export class AppModule { }

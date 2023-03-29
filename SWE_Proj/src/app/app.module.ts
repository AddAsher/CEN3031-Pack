import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AuthService } from './Login/auth.service';
import { appRoutingModule } from './app.routing';

import { LoginComponent } from './Login'; //login page
import { HomeComponent } from './Home'; //home page 
import { RegistrationComponent } from './Registration'; //registration page

@NgModule({
  declarations: [LoginComponent,HomeComponent,RegistrationComponent],
  imports: [BrowserModule, HttpClientModule, FormsModule, appRoutingModule],
  providers: [AuthService],
  bootstrap: [LoginComponent],
})
export class AppModule { }

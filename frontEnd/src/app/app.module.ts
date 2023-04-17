import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AuthService } from './auth.service';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomePageComponent } from './home-page/home-page.component';
import { RegistrationComponent } from './registration/registration.component';
import { SlideshowComponent } from './slideshow/slideshow.component';
import { AddClubComponent } from './add-club/add-club/add-club.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomePageComponent,
    RegistrationComponent,
    SlideshowComponent,
    AddClubComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    AppRoutingModule
  ],
  providers: [AuthService],
  bootstrap: [AppComponent]
})
export class AppModule { }

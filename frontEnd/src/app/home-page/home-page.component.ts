import { Component } from '@angular/core';
import { SlideshowComponent } from '../slideshow/slideshow.component';
import { ClubListComponent } from '../club-list/club-list.component';
import { AuthService, User } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})
export class HomePageComponent {

  //variables to be used to construct a search bar
  searchTerm: string;
  searchResults: any[];
  welcomeMessage: string;
  currUser: string;

  constructor(private authService: AuthService, private router: Router) {
    this.searchTerm = "";
    this.searchResults = [];
    this.welcomeMessage = "Find the club that's right for you!";
    this.currUser = "Admin"
  }


  performSearch() {
    //1. retrieve list of clubs stored in backend to be put into an array
    //2. search that array for the name if it is present
    //3. if found go from there
    //if not create a popup saying 'Club not yet in database'

  }

  logout() {
    alert("You have successfully logged out")
    this.router.navigate(['']);
  }

  getUser() {
    this.authService.getUser().subscribe(
      (response: User) => {
        this.currUser = response.username;
        console.log(response);
      },
      (error) => {
        console.log(error);
      }
    )
  }
}

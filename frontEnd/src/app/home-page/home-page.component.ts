import { Component } from '@angular/core';
import { SlideshowComponent } from '../slideshow/slideshow.component';

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

  constructor() {
    this.searchTerm = "";
    this.searchResults = [];
    this.welcomeMessage =  "Find the club that's right for you!";
  }

  performSearch() {
    //1. retrieve list of clubs stored in backend to be put into an array
    //2. search that array for the name if it is present
    //3. if found go from there
         //if not create a popup saying 'Club not yet in database'
    throw new Error('Method not implemented.');
  }
}

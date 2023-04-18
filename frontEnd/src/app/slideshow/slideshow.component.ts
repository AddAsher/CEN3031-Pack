import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-slideshow',
  templateUrl: './slideshow.component.html',
  styleUrls: ['./slideshow.component.css']
})
export class SlideshowComponent {
  @Input()
  clubName: string = "";

  showNavigation = false;
  images = [
    'assets/GoGators.jpg',
    'assets/box_logo.jpg'
  ];

  //I'll have to create more arrays to store images for each club (10 total)

  currentSlideIndex = 0;

  previousSlide() {
    this.currentSlideIndex--;
  }

  nextSlide() {
    this.currentSlideIndex++;
  }

  likeClub() {

  }
}

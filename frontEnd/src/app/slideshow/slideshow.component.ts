import { Component } from '@angular/core';

@Component({
  selector: 'app-slideshow',
  templateUrl: './slideshow.component.html',
  styleUrls: ['./slideshow.component.css']
})
export class SlideshowComponent {
  showNavigation = false;
  images = [
    'assets/GoGators.jpg',
    'assets/box_logo.jpg'
  ];

  currentSlideIndex = 0;

  previousSlide() {
    this.currentSlideIndex--;
  }

  nextSlide() {
    this.currentSlideIndex++;
  }
}

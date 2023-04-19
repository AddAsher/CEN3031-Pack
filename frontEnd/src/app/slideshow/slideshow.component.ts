import { Component, Input } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-slideshow',
  templateUrl: './slideshow.component.html',
  styleUrls: ['./slideshow.component.css']
})
export class SlideshowComponent {
  @Input()
  clubName: string = "";

  //example slideshow images
  images = [
    'assets/GoGators.jpg',
    'assets/box_logo.jpg'
  ];
  images1=[];
  images2=[];
  images3=[];
  images4=[];
  images5=[];
  images6=[];
  images7=[];
  images8=[];
  images9=[];

  constructor(private authService: AuthService, private router: Router) { }
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

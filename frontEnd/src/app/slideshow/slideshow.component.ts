import { Component, Input } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-slideshow',
  templateUrl: './slideshow.component.html',
  styleUrls: ['./slideshow.component.css']
})
export class SlideshowComponent {
  clubName: string;

  //example slideshow images
  images = ['assets/GoGators.jpg','assets/box_logo.jpg'];

  //actual images
  images1=['assets/folder1/AdoptedStudentOrganization1.jpg'];
  images2=['assets/folder2/AReasonToGive1.jpg','assets/folder2/AReasonToGive2.jpg','assets/folder2/AReasonToGive3.jpg'];
  images3=['assets/folder3/AdventChristianFellowship1.jpg','assets/folder3/AdventChristianFellowship2.jpg','assets/folder3/AdventChristianFellowship3.jpg'];
  images4=['assets/folder4/AdSociety1.jpg','assets/folder4/AdSociety.jpg','assets/folder4/AdSociety3.jpg'];
  images5=[''];
  images6=['assets/folder6/AlphaOmega1.jpg','assets/folder6/AlphaOmega1.jpg','assets/folder6/AlphaOmega1.jpg'];
  images7=['assets/folder7/AquaticAnimalHealth1.jpg','assets/folder7/AquaticAnimalHealth2.jpg','assets/folder7/AquaticAnimalHealth3.jpg'];
  images8=['assets/folder8/SocietyOfPCBuilding1.jpg','assets/folder8/SocietyOfPCBuilding2.jpg','assets/folder8/SocietyOfPCBuilding3.jpg'];
  images9=['assets/folder9/3DPrintingClub1.png','assets/folder9/3DPrintingClub2.png','assets/folder9/3DPrintingClub3.png'];

  constructor(private authService: AuthService, private router: Router) { 
    this.clubName = "";
  }

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

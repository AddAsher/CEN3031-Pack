import { Component, Input } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { SharedService } from '../shared.service';

@Component({
  selector: 'app-slideshow',
  templateUrl: './slideshow.component.html',
  styleUrls: ['./slideshow.component.css']
})
export class SlideshowComponent {

  //example slideshow images
  images: string[] = [];

  //actual images
  images1=['assets/folder1/AdoptedStudentOrganization1.jpg','assets/GoGators.jpg'];
  images2=['assets/folder2/AReasonToGive1.jpg','assets/folder2/AReasonToGive2.jpg','assets/folder2/AReasonToGive3.jpg'];
  images3=['assets/folder3/AdventChristianFellowship1.jpg','assets/folder3/AdventChristianFellowship2.jpg','assets/folder3/AdventChristianFellowship3.jpg'];
  images4=['assets/folder4/AdSociety1.jpg','assets/folder4/AdSociety.jpg','assets/folder4/AdSociety3.jpg'];
  images5=['assets/folder6/AlphaOmega1.jpg','assets/folder6/AlphaOmega1.jpg','assets/folder6/AlphaOmega1.jpg'];
  images6=['assets/folder6/AlphaOmega1.jpg','assets/folder6/AlphaOmega1.jpg','assets/folder6/AlphaOmega1.jpg'];
  images7=['assets/folder7/AquaticAnimalHealth1.jpg','assets/folder7/AquaticAnimalHealth2.jpg','assets/folder7/AquaticAnimalHealth3.jpg'];
  images8=['assets/folder8/SocietyOfPCBuilding1.jpg','assets/folder8/SocietyOfPCBuilding2.jpg','assets/folder8/SocietyOfPCBuilding3.jpg'];
  images9=['assets/folder9/3DPrintingClub1.png','assets/folder9/3DPrintingClub2.png','assets/folder9/3DPrintingClub3.png'];

  club: [string,any];

  constructor(private sharedService: SharedService){
    this.club=["",null];
  }

  ngOnInit(){
    this.sharedService.getSelectedPair().subscribe(pair => {
      this.club = pair;
    });
    this.ImageCheck();
  }

  ImageCheck(){
    if(this.club[0] == "Adopted Student Organization"){
      this.images=this.images1;
    }else if(this.club[0]=="fooclub"){
      this.images=this.images2;
    }else if(this.club[0]=="A Reason to Give"){
      this.images=this.images3;
    }else if(this.club[0]=="Adventist Christian Fellowship"){
      this.images=this.images4;
    }else if(this.club[0]=="Advertising Society"){
      this.images=this.images5;
    }else if(this.club[0]=="Alpha Omega - Jewish Dental Society"){
      this.images=this.images6;
    }else if(this.club[0]=="Aquatic Animal Health Club"){
      this.images=this.images7;
    }else if(this.club[0]=="The Society of PC Building"){
      this.images=this.images8;
    }else if(this.club[0]=="3D Printing Club"){
      this.images=this.images9;
    }
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

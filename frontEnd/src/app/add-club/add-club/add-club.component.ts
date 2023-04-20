import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-club',
  templateUrl: './add-club.component.html',
  styleUrls: ['./add-club.component.css']
})
export class AddClubComponent {

  constructor(private router: Router){}

  goToHomePage() {
    this.router.navigate(['/home']);
  }
  SubmitInfo() {
    throw new Error('Method not implemented.');
  }

}

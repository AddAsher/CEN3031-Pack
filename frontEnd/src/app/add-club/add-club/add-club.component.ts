import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-club',
  templateUrl: './add-club.component.html',
  styleUrls: ['./add-club.component.css']
})
export class AddClubComponent {
  constructor(private router: Router){}

  BackHome() {
    this.router.navigate(['/home']);
  }
  Logout() {
    this.router.navigate(['']);
  }

}

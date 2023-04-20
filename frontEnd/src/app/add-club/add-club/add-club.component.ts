import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/auth.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-add-club',
  templateUrl: './add-club.component.html',
  styleUrls: ['./add-club.component.css']
})
export class AddClubComponent {
  name: string;
  description: string = "";
  contact: string = "";
  leader: string = "";
  hyperlink: string = "";

  constructor(private authService: AuthService, private router: Router) {
    this.name = "";
  }
  goToHomePage() {
    this.router.navigate(['/home']);
  }
  SubmitInfo() {
    console.log("Adding Club:", this.name, this.description, this.contact, this.leader, this.hyperlink);
    this.authService.addClub(this.name, this.description, this.contact, this.leader, this.hyperlink).subscribe(
      (response) => {
        console.log("Request request received on Front end");
        this.goToHomePage();
        document.forms[0].reset()
      },
      (error) => {
        document.forms[0].reset()
        console.log(error);
        const errorMessage = error.error.split('\n')[0];
        alert(errorMessage);
      }
    )
  }

}

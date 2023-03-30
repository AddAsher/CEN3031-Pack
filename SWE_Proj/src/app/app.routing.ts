import { Routes, RouterModule } from '@angular/router';

import { LoginComponent } from './Login';
import { HomeComponent } from './Home';
import { RegistrationComponent } from './Registration';

const routes: Routes = [
    { path: 'login', component: LoginComponent },
    { path: 'home', component: HomeComponent},
    { path: 'register', component: RegistrationComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: 'login' }
];

export const appRoutingModule = RouterModule.forRoot(routes);
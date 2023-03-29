import { Routes, RouterModule } from '@angular/router';

import { HomeComponent } from './Home';
import { RegistrationComponent } from './Registration';

const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'register', component: RegistrationComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];

export const appRoutingModule = RouterModule.forRoot(routes);
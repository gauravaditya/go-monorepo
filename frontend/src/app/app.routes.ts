
import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home-page/home.component';
import { CreateEventComponent } from './pages/create-events/create-event.component';
import { ConsumedEventsComponent } from './pages/consumed-events/consumed-events.component';

export const routes: Routes = [
	{ path: '', component: HomeComponent },
	{ path: 'create', component: CreateEventComponent },
	{ path: 'consumed', component: ConsumedEventsComponent },
	{ path: '**', redirectTo: '' }
];

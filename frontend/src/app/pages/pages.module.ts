import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { HomeComponent } from './home-page/home.component';
import { CreateEventComponent } from './create-events/create-event.component';
import { ConsumedEventsComponent } from './consumed-events/consumed-events.component';
import { RouterModule } from '@angular/router';

const components: any[] = [
    HomeComponent, 
    CreateEventComponent, 
    ConsumedEventsComponent
];
    
@NgModule({
  imports: [
    CommonModule, 
    ...components],
  exports: [...components],
})
export class PagesModule {}





import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Event } from '../../api/event.model';
import { RouterModule, Router, NavigationEnd } from '@angular/router';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-consumed-events',
  templateUrl: './consumed-events.component.html',
  styleUrls: ['./consumed-events.component.scss'],
  imports: [CommonModule, RouterModule, MatTableModule, MatButtonModule]
})
export class ConsumedEventsComponent implements OnInit {
  events: Event[] = [];

  constructor(private http: HttpClient, private router: Router) {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.loadEvents();
      }
    });
  }

  ngOnInit() {
    this.loadEvents();
  }

  loadEvents() {
    this.http.get<{events: Event[]}>('http://localhost:8080/events-data')
      .subscribe(data => {
        console.log('Fetched events:', data);
        this.events = data.events || [];
      });
  }
}
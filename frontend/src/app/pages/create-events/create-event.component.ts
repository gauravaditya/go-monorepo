

import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-create-event',
  templateUrl: './create-event.component.html',
  styleUrls: ['./create-event.component.scss'],
  imports: [RouterModule, FormsModule, MatButtonModule, MatFormFieldModule, MatInputModule]
})
export class CreateEventComponent {
  count = 0;

  constructor(private http: HttpClient) {}

  createEvent() {
    this.http.post('http://localhost:8080/register', { count: this.count })
      .subscribe({
        next: () => {
          this.count = 0;
          // Optionally show success message
        },
        error: err => {
          // Optionally show error message
        }
      });
  }
}

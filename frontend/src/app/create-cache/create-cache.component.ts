import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CacheService } from '../services/cache.service';
import { CommonModule } from '@angular/common';

@Component({
  standalone: true,
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  selector: 'app-create-cache',
  templateUrl: './create-cache.component.html',
  styleUrl:'./create-cache.component.scss'
})
export class CreateCacheComponent implements OnInit {
  @Output() valueAdded = new EventEmitter()
  cacheForm!: FormGroup;

  constructor(private cacheService: CacheService, private fb: FormBuilder) { }

  ngOnInit(): void {
    this.initializeForm();
  }

  initializeForm(): void {
    this.cacheForm = this.fb.group({
      newKey: ['', Validators.required],
      newValue: ['', Validators.required],
      expirationDate: ['', [Validators.required, Validators.pattern(/^\d+$/), Validators.min(2)]]
    });
  }

  addKeyValuePair(): void {
    if (this.cacheForm.valid) {
      const { newKey, newValue, expirationDate } = this.cacheForm.value;
      const expiresIn = parseInt(expirationDate, 10);

      this.cacheService.addValue(newKey, newValue, expiresIn).subscribe({
        next: (data) => {
          this.valueAdded.emit()
          console.log(data);
          this.clearForm();
        },
        error: (error) => {
          console.log(error);
        }
      });
    }
  }

  clearForm(): void {
    this.cacheForm.reset({
      expirationDate: '10'
    });
  }
}

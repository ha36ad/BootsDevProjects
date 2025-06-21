import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns
import tkinter as tk
from tkinter import messagebox

class ClickHeatmap:
    """
    A simple class to record mouse clicks and generate a heatmap.
    """
    
    def __init__(self, width=800, height=600, grid_size=20):
        """
        Initialize the ClickHeatmap.
        
        Args:
            width: Canvas width in pixels
            height: Canvas height in pixels
            grid_size: Size of each grid cell for the heatmap
        """
        self.width = width
        self.height = height
        self.grid_size = grid_size
        self.clicks = []
        
        # Calculate grid dimensions
        self.grid_width = width // grid_size
        self.grid_height = height // grid_size
        
        # Initialize click grid
        self.click_grid = np.zeros((self.grid_height, self.grid_width))
        
        # GUI components
        self.root = None
        self.canvas = None
        
    def add_click(self, x, y):
        """
        Add a click at the given coordinates.
        
        Args:
            x: X coordinate
            y: Y coordinate
        """
        if 0 <= x < self.width and 0 <= y < self.height:
            self.clicks.append((x, y))
            
            # Update grid
            grid_x = min(x // self.grid_size, self.grid_width - 1)
            grid_y = min(y // self.grid_size, self.grid_height - 1)
            self.click_grid[grid_y, grid_x] += 1
            
    def on_canvas_click(self, event):
        """Handle canvas click events."""
        self.add_click(event.x, event.y)
        
        # Draw a small dot where clicked
        self.canvas.create_oval(
            event.x - 2, event.y - 2, 
            event.x + 2, event.y + 2, 
            fill='red', outline='red'
        )
        
        # Update click counter
        self.update_status()
    
    def update_status(self):
        """Update the status label with click count."""
        if hasattr(self, 'status_label'):
            self.status_label.config(text=f"Clicks recorded: {len(self.clicks)}")
    
    def clear_clicks(self):
        """Clear all recorded clicks."""
        self.clicks = []
        self.click_grid = np.zeros((self.grid_height, self.grid_width))
        if self.canvas:
            self.canvas.delete("all")
        self.update_status()
    
    def show_heatmap(self):
        """Display the heatmap in a new window."""
        if len(self.clicks) == 0:
            messagebox.showwarning("No Data", "No clicks recorded yet!")
            return
        
        # Create heatmap plot
        plt.figure(figsize=(12, 8))
        
        # Use seaborn for a prettier heatmap
        sns.heatmap(
            self.click_grid, 
            cmap='YlOrRd',  # Yellow-Orange-Red colormap
            annot=False,    # Don't show numbers in cells
            fmt='d',        # Integer format if annotations were shown
            cbar_kws={'label': 'Number of Clicks'},
            square=False,   # Don't force square cells
            linewidths=0.1, # Thin lines between cells
            linecolor='white'
        )
        
        plt.title(f'Click Heatmap ({len(self.clicks)} total clicks)', fontsize=16, pad=20)
        plt.xlabel(f'Grid X (each cell = {self.grid_size} pixels)', fontsize=12)
        plt.ylabel(f'Grid Y (each cell = {self.grid_size} pixels)', fontsize=12)
        
        # Invert y-axis to match screen coordinates (seaborn does this automatically)
        plt.gca().invert_yaxis()
        
        plt.tight_layout()
        plt.show()
    
    def start_recording(self):
        """Start the GUI for recording clicks."""
        self.root = tk.Tk()
        self.root.title("Click Heatmap Recorder")
        self.root.resizable(False, False)
        
        # Create main frame
        main_frame = tk.Frame(self.root)
        main_frame.pack(padx=10, pady=10)
        
        # Instructions
        instructions = tk.Label(
            main_frame, 
            text="Click anywhere on the canvas below to record clicks.\nWhen done, click 'Show Heatmap' to see the results.",
            font=('Arial', 12)
        )
        instructions.pack(pady=(0, 10))
        
        # Canvas for clicking
        self.canvas = tk.Canvas(
            main_frame, 
            width=self.width, 
            height=self.height, 
            bg='lightgray',
            relief='sunken',
            bd=2
        )
        self.canvas.pack(pady=(0, 10))
        self.canvas.bind("<Button-1>", self.on_canvas_click)
        
        # Status and control frame
        control_frame = tk.Frame(main_frame)
        control_frame.pack(fill='x')
        
        # Status label
        self.status_label = tk.Label(control_frame, text="Clicks recorded: 0")
        self.status_label.pack(side='left')
        
        # Buttons frame
        button_frame = tk.Frame(control_frame)
        button_frame.pack(side='right')
        
        # Clear button
        clear_btn = tk.Button(
            button_frame, 
            text="Clear", 
            command=self.clear_clicks,
            bg='orange'
        )
        clear_btn.pack(side='left', padx=(0, 5))
        
        # Show heatmap button
        heatmap_btn = tk.Button(
            button_frame, 
            text="Show Heatmap", 
            command=self.show_heatmap,
            bg='lightblue'
        )
        heatmap_btn.pack(side='left', padx=(0, 5))
        
        # Quit button
        quit_btn = tk.Button(
            button_frame, 
            text="Quit", 
            command=self.root.quit,
            bg='lightcoral'
        )
        quit_btn.pack(side='left')
        
        # Start the GUI
        self.root.mainloop()
    
    def get_statistics(self):
        """Get basic statistics about the clicks."""
        if not self.clicks:
            return {"total_clicks": 0}
        
        x_coords = [click[0] for click in self.clicks]
        y_coords = [click[1] for click in self.clicks]
        
        return {
            "total_clicks": len(self.clicks),
            "x_range": (min(x_coords), max(x_coords)),
            "y_range": (min(y_coords), max(y_coords)),
            "avg_x": np.mean(x_coords),
            "avg_y": np.mean(y_coords),
            "max_clicks_in_cell": np.max(self.click_grid)
        }


class ClickSimulator:
    """
    A helper class to simulate clicks for demonstration purposes.
    """
    
    def __init__(self, heatmap):
        self.heatmap = heatmap
    
    def simulate_random_clicks(self, num_clicks=100):
        """Simulate random clicks."""
        import random
        
        for _ in range(num_clicks):
            x = random.randint(0, self.heatmap.width - 1)
            y = random.randint(0, self.heatmap.height - 1)
            self.heatmap.add_click(x, y)
    
    def simulate_hotspot_clicks(self, num_clicks=200):
        """Simulate clicks concentrated in certain areas."""
        import random
        
        # Define hotspots
        hotspots = [
            (self.heatmap.width // 4, self.heatmap.height // 4),
            (3 * self.heatmap.width // 4, self.heatmap.height // 4),
            (self.heatmap.width // 2, 3 * self.heatmap.height // 4)
        ]
        
        for _ in range(num_clicks):
            # Choose a hotspot
            hotspot = random.choice(hotspots)
            
            # Add some randomness around the hotspot
            x = max(0, min(self.heatmap.width - 1, 
                          hotspot[0] + random.randint(-50, 50)))
            y = max(0, min(self.heatmap.height - 1, 
                          hotspot[1] + random.randint(-50, 50)))
            
            self.heatmap.add_click(x, y)


def main():
    """
    Main function to run the click heatmap application.
    """
    print("Simple Click Heatmap Generator")
    print("=" * 35)
    
    # Create heatmap instance
    heatmap = ClickHeatmap(width=800, height=600, grid_size=25)
    
    # Ask user what they want to do
    print("\nChoose an option:")
    print("1. Record clicks manually (opens GUI)")
    print("2. Generate sample data and show heatmap")
    
    choice = input("Enter choice (1 or 2): ").strip()
    
    if choice == "1":
        print("\nStarting click recorder...")
        print("A window will open where you can click to record data.")
        heatmap.start_recording()
        
    elif choice == "2":
        print("\nGenerating sample click data...")
        
        # Create simulator and generate data
        simulator = ClickSimulator(heatmap)
        simulator.simulate_hotspot_clicks(150)
        simulator.simulate_random_clicks(50)
        
        # Show statistics
        stats = heatmap.get_statistics()
        print(f"\nGenerated {stats['total_clicks']} clicks")
        print(f"Average position: ({stats['avg_x']:.1f}, {stats['avg_y']:.1f})")
        print(f"Max clicks in single cell: {stats['max_clicks_in_cell']}")
        
        # Show heatmap
        print("\nDisplaying heatmap...")
        heatmap.show_heatmap()
        
    else:
        print("Invalid choice. Run the program again.")


if __name__ == "__main__":
    main()
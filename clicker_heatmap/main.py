from heatmap import ClickHeatmap, ClickSimulator
import sys

def print_banner():
    print("=" * 50)
    print("     CLICK HEATMAP GENERATOR")
    print("=" * 50)
    print("Track mouse clicks and visualize them as heatmaps")
    print()

def show_menu():
    print("Choose an option:")
    print("1. Record clicks manually (Interactive GUI)")
    print("2. Generate sample data and show heatmap")
    print("3. Exit")
    print()

def get_user_choice():
    while True:
        try:
            choice = input("Enter your choice (1-3): ").strip()
            if choice in ['1', '2', '3']:
                return int(choice)
            else:
                print("Please enter a number between 1 and 5.")
        except KeyboardInterrupt:
            print("\nGoodbye!")
            sys.exit(0)
        except Exception:
            print("Invalid input. Please try again.")

#Start manual click recording with default settings.
def manual_recording():
    print("\n Starting manual click recording...")
    print("A window will open where you can click to record data.")
    print("Click anywhere on the gray canvas to record clicks.")
    print("Use 'Show Heatmap' button to visualize your clicks.")
    print()
    
    heatmap = ClickHeatmap(width=800, height=600, grid_size=25)
    heatmap.start_recording()
    
    # Show final statistics after GUI closes
    stats = heatmap.get_statistics()
    if stats['total_clicks'] > 0:
        print(f"\n Final Statistics:")
        print(f"   Total clicks: {stats['total_clicks']}")
        print(f"   Average position: ({stats['avg_x']:.1f}, {stats['avg_y']:.1f})")
        print(f"   Max clicks in single cell: {stats['max_clicks_in_cell']}")

def sample_data_demo():
    print("\nGenerating sample click data...")
    
    # Create heatmap with default settings
    heatmap = ClickHeatmap(width=800, height=600, grid_size=20)
    
    # Generate sample data
    simulator = ClickSimulator(heatmap)
    
    print("   Creating hotspots with concentrated clicks...")
    simulator.simulate_hotspot_clicks(200)
    
    print("   Adding random background clicks...")
    simulator.simulate_random_clicks(80)
    
    # Show statistics
    stats = heatmap.get_statistics()
    print(f"\nGenerated Statistics:")
    print(f"   Total clicks: {stats['total_clicks']}")
    print(f"   Coverage area: {stats['x_range'][1] - stats['x_range'][0]} x {stats['y_range'][1] - stats['y_range'][0]} pixels")
    print(f"   Average position: ({stats['avg_x']:.1f}, {stats['avg_y']:.1f})")
    print(f"   Max clicks in single cell: {stats['max_clicks_in_cell']}")
    
    print("\n Displaying heatmap...")
    heatmap.show_heatmap()

def main():
    try:
        print_banner()
        while True:
            show_menu()
            choice = get_user_choice()
            if choice == 1:
                manual_recording()
            elif choice == 2:
                sample_data_demo()
            elif choice == 3:
                print("Thank you for using Click Heatmap Generator!")
                break
            if choice != 3:
                print("\n" + "─" * 50)
                continue_choice = input("Press Enter to return to main menu (or 'q' to quit): ").strip().lower()
                if continue_choice == 'q':
                    print("\n Goodbye!")
                    break
                print()
    
    except KeyboardInterrupt:
        print("\n\n Goodbye!")
    except Exception as e:
        print(f"\n An unexpected error occurred: {e}")
        print("Please check that all required libraries are installed:")
        print(" pip install numpy matplotlib seaborn")

if __name__ == "__main__":
    main()
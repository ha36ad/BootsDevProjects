from heatmap import ClickHeatmap, ClickSimulator
import sys

def print_banner():
    """Print application banner."""
    print("=" * 50)
    print("     CLICK HEATMAP GENERATOR")
    print("=" * 50)
    print("Track mouse clicks and visualize them as heatmaps")
    print()

def show_menu():
    """Display the main menu options."""
    print("Choose an option:")
    print("1. 🖱️  Record clicks manually (Interactive GUI)")
    print("2. 📊 Generate sample data and show heatmap")
    print("3. ⚙️  Custom configuration")
    print("4. ❓ Help")
    print("5. 🚪 Exit")
    print()

def get_user_choice():
    """Get and validate user input."""
    while True:
        try:
            choice = input("Enter your choice (1-5): ").strip()
            if choice in ['1', '2', '3', '4', '5']:
                return int(choice)
            else:
                print("❌ Please enter a number between 1 and 5.")
        except KeyboardInterrupt:
            print("\n\n👋 Goodbye!")
            sys.exit(0)
        except Exception:
            print("❌ Invalid input. Please try again.")

def manual_recording():
    """Start manual click recording with default settings."""
    print("\n🖱️  Starting manual click recording...")
    print("📝 A window will open where you can click to record data.")
    print("💡 Click anywhere on the gray canvas to record clicks.")
    print("🎨 Use 'Show Heatmap' button to visualize your clicks.")
    print()
    
    heatmap = ClickHeatmap(width=800, height=600, grid_size=25)
    heatmap.start_recording()
    
    # Show final statistics after GUI closes
    stats = heatmap.get_statistics()
    if stats['total_clicks'] > 0:
        print(f"\n📈 Final Statistics:")
        print(f"   Total clicks: {stats['total_clicks']}")
        print(f"   Average position: ({stats['avg_x']:.1f}, {stats['avg_y']:.1f})")
        print(f"   Max clicks in single cell: {stats['max_clicks_in_cell']}")

def sample_data_demo():
    """Generate and display sample data."""
    print("\n📊 Generating sample click data...")
    
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
    print(f"\n📈 Generated Statistics:")
    print(f"   Total clicks: {stats['total_clicks']}")
    print(f"   Coverage area: {stats['x_range'][1] - stats['x_range'][0]} x {stats['y_range'][1] - stats['y_range'][0]} pixels")
    print(f"   Average position: ({stats['avg_x']:.1f}, {stats['avg_y']:.1f})")
    print(f"   Max clicks in single cell: {stats['max_clicks_in_cell']}")
    
    print("\n🎨 Displaying heatmap...")
    heatmap.show_heatmap()

def custom_configuration():
    """Allow user to customize heatmap settings."""
    print("\n⚙️  Custom Configuration")
    print("Configure your heatmap settings:")
    
    try:
        # Get custom dimensions
        print("\n📐 Canvas Size:")
        width = int(input(f"   Width in pixels (default 800): ") or "800")
        height = int(input(f"   Height in pixels (default 600): ") or "600")
        
        # Get grid size
        print("\n🔲 Grid Settings:")
        grid_size = int(input(f"   Grid cell size in pixels (default 25): ") or "25")
        
        # Create custom heatmap
        heatmap = ClickHeatmap(width=width, height=height, grid_size=grid_size)
        
        print(f"\n✅ Created custom heatmap: {width}x{height} pixels, {grid_size}px grid")
        
        # Ask what to do next
        print("\nWhat would you like to do?")
        print("1. Start recording clicks")
        print("2. Generate sample data")
        
        sub_choice = input("Choice (1 or 2): ").strip()
        
        if sub_choice == "1":
            print("\n🖱️  Starting custom click recording...")
            heatmap.start_recording()
        elif sub_choice == "2":
            print("\n📊 Generating sample data for custom configuration...")
            simulator = ClickSimulator(heatmap)
            simulator.simulate_hotspot_clicks(150)
            simulator.simulate_random_clicks(50)
            
            stats = heatmap.get_statistics()
            print(f"Generated {stats['total_clicks']} clicks")
            heatmap.show_heatmap()
        else:
            print("❌ Invalid choice, returning to main menu.")
            
    except ValueError:
        print("❌ Invalid input. Please enter valid numbers.")
    except Exception as e:
        print(f"❌ Error: {e}")

def show_help():
    """Display help information."""
    print("\n❓ HELP - Click Heatmap Generator")
    print("=" * 40)
    print()
    print("📖 What is this?")
    print("   This tool helps you visualize mouse click patterns by creating")
    print("   colorful heatmaps showing where clicks are concentrated.")
    print()
    print("🎯 How to use:")
    print("   1. Choose 'Record clicks manually' to open an interactive window")
    print("   2. Click anywhere on the gray canvas to record click positions")
    print("   3. Press 'Show Heatmap' to see your click pattern visualization")
    print("   4. Red/orange areas = high click concentration")
    print("   5. Yellow/white areas = low click concentration")
    print()
    print("⚙️  Custom Configuration:")
    print("   - Canvas Size: Larger = more detailed tracking area")
    print("   - Grid Size: Smaller = higher resolution heatmap")
    print()
    print("💡 Tips:")
    print("   - Try clicking in patterns to see clear hotspots")
    print("   - Use 'Clear' button to reset and start over")
    print("   - Sample data shows typical user behavior patterns")
    print()
    print("🔧 Technical:")
    print("   - Grid cells count clicks within each area")
    print("   - Heatmap uses color intensity to show click density")
    print("   - Seaborn library provides professional visualization")
    print()

def main():
    """Main application entry point."""
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
                custom_configuration()
            elif choice == 4:
                show_help()
            elif choice == 5:
                print("\n👋 Thank you for using Click Heatmap Generator!")
                print("💡 Remember: Understanding user behavior helps create better interfaces!")
                break
            
            # Ask if user wants to continue
            if choice != 5:
                print("\n" + "─" * 50)
                continue_choice = input("Press Enter to return to main menu (or 'q' to quit): ").strip().lower()
                if continue_choice == 'q':
                    print("\n👋 Goodbye!")
                    break
                print()
    
    except KeyboardInterrupt:
        print("\n\n👋 Goodbye!")
    except Exception as e:
        print(f"\n❌ An unexpected error occurred: {e}")
        print("Please check that all required libraries are installed:")
        print("   pip install numpy matplotlib seaborn")

if __name__ == "__main__":
    main()
from Plotter import Plotter
import sys, getopt

def main(args):
    path = ''
    try:
        opts, args = getopt.getopt(args,"f:",["file="])
    except getopt.GetoptError:
        print("graph.py -f <foldername>")
        sys.exit(2)
    for opt, arg in opts:
        if opt in ('-f', '--file'):
            path = arg
    print("Input file:", path)
    plotter = Plotter(path)
    plotter.plot()

if __name__ == "__main__":
    main(sys.argv[1:])

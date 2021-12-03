day = $(shell date +'%-d')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
		mkdir calendar/day-0$(day)/pt1; \
		mkdir calendar/day-0$(day)/pt2; \
  		cp template calendar/day-0$(day)/pt1/main.go; \
  		cp template calendar/day-0$(day)/pt2/main.go; \
  	else \
  		mkdir calendar/day-$(day); \
		mkdir calendar/day-$(day)/pt1; \
		mkdir calendar/day-$(day)/pt2; \
		cp template calendar/day-$(day)/pt1/main.go; \
		cp template calendar/day-$(day)/pt2/main.go; \
    fi
	@echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/


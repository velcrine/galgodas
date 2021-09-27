# galgodas
Real problems based implementations of standard algorithms 
and data structures in golang, along with their variations, 
and decent lectures before taking each step

## inspiration
this effort was made under a workshop instigated by mathematician @mohit-baliyan 
to init the confidence among his CS students for the ability to render every 
standard ds and frequently used algorithms in general life coding.

## as a result
We have implemented several packages which could be called from other 
programs and perform those tasks.

    For ex.
    Selection sort seems pretty useless.
    But if the question statement is : 
    
    Out a group of students, I want only the smallest one,
    then the smallest one left among all : quick sort will seem pretty useless.

    As selection sort will give the shortest of all in "n" time complexity one by one.
    QuickSort will have to sort all the elements with NlogN, and then will give out just 1.
    
    However, output out of a selection sort program is not that easy to gain,
    Proper care has been taken for usability,
    Hence, the output comes through a golang channel. Oncee channel is consumed, then only 
    the next value is calculated.

